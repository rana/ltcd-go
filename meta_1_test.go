package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

// Configuration struct for the application
type Config struct {
	Hosts          []string
	ProcessName    string
	SSHUser        string
	SSHKeyPath     string
	MaxConcurrent  int
	EmailFrom      string
	EmailTo        []string
	SMTPServer     string
	SMTPPort       int
	SMTPUser       string
	SMTPPassword   string
	ConnectionTime time.Duration
}

// Result of process checking on a single host
type HostResult struct {
	Host    string
	Success bool
	Message string
	Running bool
}

// LoadConfig loads application configuration
func LoadConfig() (*Config, error) {
	// In a real application, this would load from a config file or environment
	// This is a simplified example with hardcoded values
	return &Config{
		Hosts:          generateHosts(100), // Generate 100 example hosts
		ProcessName:    "target_process",
		SSHUser:        "admin",
		SSHKeyPath:     os.Getenv("HOME") + "/.ssh/id_rsa",
		MaxConcurrent:  20, // Check 20 hosts concurrently
		EmailFrom:      "monitor@example.com",
		EmailTo:        []string{"admin@example.com"},
		SMTPServer:     "smtp.example.com",
		SMTPPort:       587,
		SMTPUser:       "monitor@example.com",
		SMTPPassword:   "password",
		ConnectionTime: 10 * time.Second,
	}, nil
}

// generateHosts creates a list of example host addresses
func generateHosts(cnt int) []string {
	hst := make([]string, cnt)
	for i := range cnt {
		hst[i] = fmt.Sprintf("server-%03d.example.com", i+1)
	}
	return hst
}

// CheckHost connects to a host and checks if the process is running
func CheckHost(cfg *Config, hst string) HostResult {
	rslt := HostResult{
		Host:    hst,
		Success: false,
		Running: false,
	}

	// Read private key
	key, err := os.ReadFile(cfg.SSHKeyPath)
	if err != nil {
		rslt.Message = fmt.Sprintf("Failed to read SSH key: %v", err)
		return rslt
	}

	// Create signer
	sgn, err := ssh.ParsePrivateKey(key)
	if err != nil {
		rslt.Message = fmt.Sprintf("Failed to parse SSH key: %v", err)
		return rslt
	}

	// Configure SSH client
	sshCfg := &ssh.ClientConfig{
		User: cfg.SSHUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(sgn),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Not secure for production
		Timeout:         cfg.ConnectionTime,
	}

	// Connect to host
	cln, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", hst), sshCfg)
	if err != nil {
		rslt.Message = fmt.Sprintf("Failed to connect: %v", err)
		return rslt
	}
	defer cln.Close()

	// Create session
	ssn, err := cln.NewSession()
	if err != nil {
		rslt.Message = fmt.Sprintf("Failed to create session: %v", err)
		return rslt
	}
	defer ssn.Close()

	// Create buffer to capture output
	var out bytes.Buffer
	ssn.Stdout = &out

	// Run command to check for process
	cmd := fmt.Sprintf("pgrep -f '%s' || echo 'NOT RUNNING'", cfg.ProcessName)
	if err := ssn.Run(cmd); err != nil {
		rslt.Message = fmt.Sprintf("Command execution failed: %v", err)
		return rslt
	}

	output := strings.TrimSpace(out.String())
	rslt.Success = true

	if output == "NOT RUNNING" {
		rslt.Message = fmt.Sprintf("Process '%s' is not running", cfg.ProcessName)
		rslt.Running = false
	} else {
		rslt.Message = fmt.Sprintf("Process '%s' is running (PID: %s)", cfg.ProcessName, output)
		rslt.Running = true
	}

	return rslt
}

// generateReport creates a report from the host check results
func generateReport(cfg *Config, rslts []HostResult) string {
	var rep strings.Builder

	tstamp := time.Now().Format("2006-01-02 15:04:05")
	rep.WriteString(fmt.Sprintf("Process Monitor Report - %s\n\n", tstamp))
	rep.WriteString(fmt.Sprintf("Process Name: %s\n", cfg.ProcessName))
	rep.WriteString(fmt.Sprintf("Hosts Checked: %d\n\n", len(rslts)))

	// Count statistics
	connFail := 0
	runCnt := 0
	notRunCnt := 0

	for _, rslt := range rslts {
		if !rslt.Success {
			connFail++
		} else if rslt.Running {
			runCnt++
		} else {
			notRunCnt++
		}
	}

	rep.WriteString("SUMMARY:\n")
	rep.WriteString(fmt.Sprintf("- Connection Successful: %d\n", len(rslts)-connFail))
	rep.WriteString(fmt.Sprintf("- Connection Failed: %d\n", connFail))
	rep.WriteString(fmt.Sprintf("- Process Running: %d\n", runCnt))
	rep.WriteString(fmt.Sprintf("- Process Not Running: %d\n\n", notRunCnt))

	rep.WriteString("HOSTS WITH PROCESS RUNNING:\n")
	for _, rslt := range rslts {
		if rslt.Success && rslt.Running {
			rep.WriteString(fmt.Sprintf("- %s\n", rslt.Host))
		}
	}
	rep.WriteString("\n")

	rep.WriteString("HOSTS WITHOUT PROCESS RUNNING:\n")
	for _, rslt := range rslts {
		if rslt.Success && !rslt.Running {
			rep.WriteString(fmt.Sprintf("- %s\n", rslt.Host))
		}
	}
	rep.WriteString("\n")

	rep.WriteString("CONNECTION FAILURES:\n")
	for _, rslt := range rslts {
		if !rslt.Success {
			rep.WriteString(fmt.Sprintf("- %s: %s\n", rslt.Host, rslt.Message))
		}
	}

	return rep.String()
}

// sendEmail sends the report via email
func sendEmail(cfg *Config, rep string) error {
	// Set up authentication information
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPassword, cfg.SMTPServer)

	// Construct email
	tstamp := time.Now().Format("2006-01-02")
	sub := fmt.Sprintf("Process Monitor Report - %s", tstamp)
	to := strings.Join(cfg.EmailTo, ", ")

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"%s", to, cfg.EmailFrom, sub, rep))

	addr := fmt.Sprintf("%s:%d", cfg.SMTPServer, cfg.SMTPPort)
	return smtp.SendMail(addr, auth, cfg.EmailFrom, cfg.EmailTo, msg)
}

func TestMeta1(t *testing.T) {
	// Load configuration
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting process check for '%s' on %d hosts", cfg.ProcessName, len(cfg.Hosts))

	// Create channel for results
	rslts := make([]HostResult, 0, len(cfg.Hosts))
	rsltChn := make(chan HostResult)

	// Create semaphore to limit concurrent connections
	sem := make(chan struct{}, cfg.MaxConcurrent)

	// Start worker goroutines
	var wg sync.WaitGroup
	for _, hst := range cfg.Hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			sem <- struct{}{}        // Acquire semaphore
			defer func() { <-sem }() // Release semaphore

			rslt := CheckHost(cfg, h)
			rsltChn <- rslt
		}(hst)
	}

	// Create goroutine to close result channel when all workers done
	go func() {
		wg.Wait()
		close(rsltChn)
	}()

	// Collect results
	for rslt := range rsltChn {
		rslts = append(rslts, rslt)
		if rslt.Success {
			log.Printf("Host %s: %s", rslt.Host, rslt.Message)
		} else {
			log.Printf("Host %s: Connection failed - %s", rslt.Host, rslt.Message)
		}
	}

	// Generate report
	log.Printf("Generating report for %d hosts", len(rslts))
	rep := generateReport(cfg, rslts)

	// Send email report
	log.Printf("Sending email report to %s", strings.Join(cfg.EmailTo, ", "))
	if err := sendEmail(cfg, rep); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Process monitoring completed successfully")
}
