// Package performance measurement ability with sub-millisecond resolution.
//
// This includes:
//
// * An insights into the latency of certain events triggered by user
// interactions.
//
// * Detect presence of "long tasks" that monopolize the UI thread for
// extended periods of time and block other critical tasks from being
// executed - e.g. reacting to user input.
//
// * Complete timing information on navigation of a document.
//
// * Capture a series of key moments (First Paint, First Contentful Paint)
// during pageload which developers care about.
//
// * Let a server to communicate performance metrics about the request-response
// cycle to the user agent. It also standardizes a JavaScript interface to
// enable applications to collect, process, and act on these metrics to
// optimize application delivery.
package performance
