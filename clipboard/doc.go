// Package clipboard control the clipboard
package clipboard

// Permission name to use for reading the clipboard.
const ReadPermissionName = "clipboard-read"

// Permission name to use for writing the clipboard.
//
// Note: there is no relation ship between the permission names,
// they are according to specification, two different permissions
// with no inherent strength relationship. The browser may apply
// differet rules.
const WritePermissionname = "clipboard-write"
