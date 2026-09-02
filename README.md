# 🛏️ tuck

> A lightweight terminal session manager — detach and reattach without the complexity.

**tuck** is like `tmux` or `screen`, but simpler. No window splitting, no status bars — just session management that stays out of your way.

## 🎯 Why I Built This

I wanted to run [Claude Code](https://github.com/anthropics/claude-code) remotely via SSH from my phone while on the go. But with `tmux`, `screen`, or `abduco`:

- 📱 **Screen rendering gets messy** — Alternate screen buffer doesn't play nice with mobile SSH clients
- 📜 **Can't scroll** — Terminal's native scrollback is hijacked
- 🤯 **Too much complexity** — I just want detach/attach, not window management

**tuck** solves this by *not* using the alternate screen buffer. Your terminal stays clean, scrollback works, and Claude Code renders perfectly.

## 🤔 Why tuck?

| Feature | tmux | screen | abduco | tuck |
|---------|------|--------|--------|------|
| Session detach/attach | ✅ | ✅ | ✅ | ✅ |
| Multiple clients | ✅ | ✅ | ✅ | ✅ |
| Window splitting | ✅ | ✅ | ❌ | ❌ |
| Status bar | ✅ | ✅ | ❌ | ❌ |
| Native scrollback | ❌ | ❌ | ❌ | ✅ |
| Zero config | ❌ | ❌ | ✅ | ✅ |

**tuck** doesn't use the alternate screen buffer, so your terminal's scrollback buffer remains functional. Perfect for tools like Claude Code that rely on terminal rendering.

## ✨ Features

- **📎 Session Management** — Create, attach, detach, and delete sessions
- **👥 Session Sharing** — Multiple clients can connect to the same session (pair programming!)
- **📜 Scrollback Works** — Unlike tmux/screen, your terminal's native scrollback keeps working
- **🎯 Zero Config** — No configuration files needed
- **🪶 Lightweight** — Single binary, minimal dependencies

## 📦 Installation

### Download Binary

Download the latest release from [GitHub Releases](https://github.com/rot1024/tuck/releases).

### Go Install

```bash
go install github.com/rot1024/tuck@latest
```

### Build from Source

```bash
git clone https://github.com/rot1024/tuck.git
cd tuck
go build
```

## 🚀 Quick Start

```bash
# Start a new session (auto-generated name from current directory)
tuck
tuck new

# Start with a specific command (auto-generated name)
tuck new bash

# Start with a specific name
tuck create myproject

# Start with a specific name and command
tuck create myproject bash

# List sessions (shows name, last active time, command)
tuck list
# myproject    5s ago     claude
# dev          2h ago     bash

# Attach to an existing session
tuck attach myproject

# Attach to the most recently active session
tuck attach

# Delete a session
tuck delete myproject
```

## ⌨️ Keybindings

| Key | Action |
|-----|--------|
| `~.` | Detach from session (after Enter, like SSH) |

### Escape Sequence

You can detach by pressing `~.` (tilde then period) after a newline. This works great with Claude Code and other applications that capture control keys.

### Custom Detach Key

You can configure detach keys via flags or environment variables:

```bash
# Single key via flag
tuck -d '~.' new
tuck -d ctrl-a attach mysession

# Multiple keys via flags
tuck -d '`.' -d ctrl-a new

# Via environment variables
export TUCK_DETACH_KEY='`.'
export TUCK_DETACH_KEY_1='~.'
export TUCK_DETACH_KEY_2=ctrl-a
tuck new
```

Supported formats:
- Escape sequences: `` `. ``, `~.` (character + period, triggered after Enter)
- Control keys: `ctrl-a`, `ctrl-]`, `^a`, `^A`

## 💬 Messages

tuck shows helpful status messages:

```
[tuck: ✨ created "myproject" (~. to detach)]
[tuck: 🔗 attached "myproject" (~. to detach)]
[tuck: 👋 detached "myproject"]
[tuck: 🏁 ended "myproject"]
```

Use `--quiet` or `-q` to suppress messages.

## 📝 Commands

```
tuck                      # Create and attach to a new session (auto-named)
tuck new [cmd]            # Create a new session with auto-generated name
tuck create <name> [cmd]  # Create a new session with specified name
tuck attach [name]        # Attach to a session (default: most recent)
tuck list                 # List all sessions (with last active time)
tuck delete <name>        # Delete a session
tuck clear                # Delete all sessions
```

### Aliases

- `tuck n` → `tuck new`
- `tuck c` → `tuck create`
- `tuck a` → `tuck attach`
- `tuck ls` → `tuck list`
- `tuck rm` → `tuck delete`

## 🔍 Knowing You're Inside tuck

tuck deliberately has no status bar, but there are two lightweight, opt-in
ways to tell you're inside a session:

### Terminal Title (built-in)

Pass `-T` / `--title` (or set `TUCK_TITLE=1`) and tuck will show the session
name in your terminal's tab/window title while attached, e.g. `🛏 myproject`.
This uses the standard OSC 2 escape sequence — no status bar, no extra
screen line, and it plays nicely with scrollback:

```bash
tuck -T attach myproject
export TUCK_TITLE=1   # enable for every tuck invocation

# customize the format ({name} is replaced with the session name)
tuck --title-format 'tuck:{name}' attach myproject
export TUCK_TITLE_FORMAT='tuck:{name}'
```

The title reverts to blank on detach (most shells with a fancy prompt will
overwrite it again anyway).

### Shell Prompt Indicator (manual)

Since `TUCK_SESSION` is exported inside every tuck session, you can show a
small colored marker in your prompt yourself. This is the closest thing to
a tmux-style status indicator, without tuck imposing any UI:

**bash / zsh:**

```bash
if [ -n "$TUCK_SESSION" ]; then
  PS1="%F{yellow}🛏 $TUCK_SESSION%f $PS1"   # zsh
  # PS1="\[\e[33m\]🛏 $TUCK_SESSION\[\e[0m\] $PS1"   # bash
fi
```

**fish:**

```fish
function fish_prompt
  if set -q TUCK_SESSION
    set_color yellow; echo -n "🛏 $TUCK_SESSION "; set_color normal
  end
  # ...rest of your prompt
end
```

Both approaches are entirely optional and off by default, in keeping with
tuck's zero-config philosophy.

## 🔧 Environment Variables

| Variable | Description |
|----------|-------------|
| `TUCK_SESSION` | Set inside tuck sessions. Prevents nested tuck sessions. |
| `TUCK_DETACH_KEY` | Default detach key (e.g., `~.`, `` `. ``, `ctrl-a`) |
| `TUCK_DETACH_KEY_1`, `_2`, ... | Additional detach keys |
| `TUCK_TITLE` | Set to show the session name in the terminal title while attached |
| `TUCK_TITLE_FORMAT` | Custom terminal title format, `{name}` is replaced with the session name (implies `TUCK_TITLE`) |

## 📄 License

MIT
