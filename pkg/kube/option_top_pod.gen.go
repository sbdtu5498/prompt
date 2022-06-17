// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/paralus/prompt/pkg/prompt"
)

var topPodOptions = []prompt.Suggest{
	{Text: "-A", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	{Text: "--containers", Description: "If present, print usage of containers within a pod."},
	{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in"},
	{Text: "--heapster-port", Description: "Port name in service to use"},
	{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as"},
	{Text: "--heapster-service", Description: "Name of Heapster service"},
	{Text: "--no-headers", Description: "If present, print output without headers."},
	{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	{Text: "--sort-by", Description: "If non-empty, sort pods list using specified field. The field can be either 'cpu' or 'memory'."},
}
