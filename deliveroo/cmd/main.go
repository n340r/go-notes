package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Example cron expression:

// */15 0 1,15 * 1-5 /usr/bin/find

const minFields = 6

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Usage: %s <YOUR CRON EXPRESSION>", os.Args[0])
	}

	cronExpression := os.Args[1]
	minutes, hours, daysOfMonth, months, daysOfWeek, command, err := parseCronExpression(cronExpression)
	if err != nil {
		return err
	}

	output(os.Stdout, minutes, hours, daysOfMonth, months, daysOfWeek, command)

	return nil
}

func parseCronExpression(cronExpression string) (minutes, hours, daysOfMonth, months, daysOfWeek []int, command string, err error) {
	if cronExpression == "" {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("cron expression is empty")
	}

	parts := strings.Fields(cronExpression)
	if len(parts) < minFields {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("invalid cron expression: %q must be at least %d fields, e.g.: "+
			"*/15 0 1,15 * 1-5 /usr/bin/find", cronExpression, minFields)
	}

	minutes, err = parse(parts[0], 0, 59)
	if err != nil {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("parse minutes: %s", err)
	}
	hours, err = parse(parts[1], 0, 23)
	if err != nil {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("parse hours: %s", err)
	}
	daysOfMonth, err = parse(parts[2], 1, 31)
	if err != nil {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("parse days of month: %s", err)
	}
	months, err = parse(parts[3], 1, 12)
	if err != nil {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("parse months: %s", err)
	}
	daysOfWeek, err = parse(parts[4], 0, 6)
	if err != nil {
		return nil, nil, nil, nil, nil, "", fmt.Errorf("parse days of week: %s", err)
	}

	command = strings.Join(parts[5:], " ")
	return minutes, hours, daysOfMonth, months, daysOfWeek, command, nil
}

func parse(field string, min, max int) ([]int, error) {
	if field == "" {
		return nil, fmt.Errorf("empty field")
	}

	seen := make(map[int]bool)

	segments := strings.Split(field, ",")
	for _, seg := range segments {
		if seg == "" {
			return nil, fmt.Errorf("empty comma segment")
		}

		rangeStr := seg
		step := 1

		if strings.Contains(seg, "/") {
			var err error
			rangeStr, step, err = parseExpressionWithStep(seg)
			if err != nil {
				return nil, err
			}
		}

		values, err := parseExpression(rangeStr, min, max)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(values); i += step {
			seen[values[i]] = true
		}
	}

	result := make([]int, 0, len(seen))
	for v := range seen {
		result = append(result, v)
	}
	sort.Ints(result)

	return result, nil
}

func parseExpressionWithStep(s string) (string, int, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid range/step: %s", s)
	}

	step, err := strconv.Atoi(parts[1])
	if err != nil || step <= 0 {
		return "", 0, fmt.Errorf("invalid step: %s", parts[1])
	}

	if parts[0] == "" {
		return "", 0, fmt.Errorf("empty range in %s", s)
	}

	return parts[0], step, nil
}

func parseExpression(expr string, min, max int) ([]int, error) {
	if expr == "*" {
		vals := make([]int, 0, max-min+1)
		for i := min; i <= max; i++ {
			vals = append(vals, i)
		}
		return vals, nil
	}

	if !strings.Contains(expr, "-") {
		v, err := strconv.Atoi(expr)
		if err != nil || v < min || v > max {
			return nil, fmt.Errorf("invalid value: %s", expr)
		}
		return []int{v}, nil
	}

	parts := strings.Split(expr, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid range: %s", expr)
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || start > end || start < min || end > max {
		return nil, fmt.Errorf("invalid range: %s", expr)
	}

	vals := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		vals = append(vals, i)
	}
	return vals, nil
}

func output(
	w io.Writer,
	minutes, hours, daysOfMonth, months, daysOfWeek []int,
	command string,
) {
	fmt.Fprintf(w, "%-14s%s\n", "minute", strings.Join(intsToStrings(minutes), " "))
	fmt.Fprintf(w, "%-14s%s\n", "hour", strings.Join(intsToStrings(hours), " "))
	fmt.Fprintf(w, "%-14s%s\n", "day of month", strings.Join(intsToStrings(daysOfMonth), " "))
	fmt.Fprintf(w, "%-14s%s\n", "month", strings.Join(intsToStrings(months), " "))
	fmt.Fprintf(w, "%-14s%s\n", "day of week", strings.Join(intsToStrings(daysOfWeek), " "))
	fmt.Fprintf(w, "%-14s%s\n", "command", command)
}

func intsToStrings(ints []int) []string {
	strs := make([]string, 0, len(ints))
	for _, i := range ints {
		strs = append(strs, strconv.Itoa(i))
	}
	return strs
}
