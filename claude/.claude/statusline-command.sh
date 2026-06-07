#!/usr/bin/env bash
input=$(cat)
cwd=$(echo "$input" | jq -r '.workspace.current_dir // .cwd')
model=$(echo "$input" | jq -r '.model.display_name // empty')
used=$(echo "$input" | jq -r '.context_window.used_percentage // empty')

# PS1-style: user@host:cwd
printf '\033[01;32m%s@%s\033[00m:\033[01;34m%s\033[00m' "$(whoami)" "$(hostname -s)" "$cwd"

# model name
[ -n "$model" ] && printf ' \033[00;33m[%s]\033[00m' "$model"

# context usage
[ -n "$used" ] && printf ' \033[00;36mctx:%.0f%%\033[00m' "$used"
