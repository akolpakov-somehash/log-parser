package parser

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

type LogStat struct {
	IP        string
	TimeStamp string
	Method    string
	Path      string
	Status    int
	BytesSent int
	Referer   string
	UserAgent string
}

var ErrParseFailed = errors.New("failed to parse log record")

type Parser interface {
	ParseRecord(record string) (LogStat, error)
	ParseFile(file string) ([]LogStat, error)
}

type HttpParser struct {
	format *regexp.Regexp
}

// sample log format:
var defaultFormat = regexp.MustCompile(`^(?P<remote_addr>\S+) - (?P<remote_user>\S+) \[(?P<time_local>.*?)\] "(?P<method>[A-Z]+) (?P<path>\S+) (?P<protocol>[A-Z]+/\d\.\d)" (?P<status>\d{3}) (?P<body_bytes_sent>\d+) "(?P<http_referer>.*?)" "(?P<http_user_agent>.*?)"$`)

func (p HttpParser) ParseRecord(record string) (LogStat, error) {
	match := p.format.FindStringSubmatch(record)
	if match == nil {
		return LogStat{}, ErrParseFailed
	}

	result := make(map[string]string)
	for i, name := range p.format.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	status, err := strconv.Atoi(result["status"])
	if err != nil {
		return LogStat{}, err
	}
	bytesSent, err := strconv.Atoi(result["body_bytes_sent"])
	if err != nil {
		return LogStat{}, err
	}

	return LogStat{
		IP:        result["remote_addr"],
		TimeStamp: result["time_local"],
		Method:    result["method"],
		Path:      result["path"],
		Status:    status,
		BytesSent: bytesSent,
		Referer:   result["http_referer"],
		UserAgent: result["http_user_agent"],
	}, nil
}

func (p HttpParser) ProcessFile(file string, parser Parser) ([]LogStat, error) {
	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	stats := make([]LogStat, 0)
	for fileScanner.Scan() {
		record := fileScanner.Text()
		stat, err := p.ParseRecord(record)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}
	return stats, nil
}

func NewHttpParser(format string) *HttpParser {
	var reFormat *regexp.Regexp
	if format == "" {
		reFormat = defaultFormat
	} else {
		reFormat = regexp.MustCompile(format)
	}
	return &HttpParser{format: reFormat}
}
