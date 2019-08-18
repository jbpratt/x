# Autumn Themer

A commandline interface for theming

Merge new resource file into old

Goals:
- Simple, easy to use
- Allow for user to provide their config and chosen theme
- Something like: `$ autumn Xresources ocean-blue` or `$ autumn i3 {theme-here}`

- Build theme and save it
- Possible 30sec timer to swap back like a confirm if possible
- Modularly import
- Backups
- Read specific config file
  - At a specific path
    - Fallback to different paths
  - Different kinds of config files
- Find setting keys on lines with hex values
  - Multiple on one line?
  - Ambiguous keys
  - Not found settings
  - Preserving current non-theme settings
  - Replace value of key with matching from theme
- Output new config file
  - Reset/Refresh program for new config 
