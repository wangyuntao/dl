# Key bindings
# ------------
if [[ $- == *i* ]]; then

# Ensure precmds are run after cd
dl-redraw-prompt() {
  local precmd
  for precmd in $precmd_functions; do
    $precmd
  done
  zle reset-prompt
}
zle -N dl-redraw-prompt

# ALT-T - cd into the selected directory
dl-cd-widget() {
  setopt localoptions pipefail 2> /dev/null
  local dir="$(dl)"
  if [[ -z "$dir" ]]; then
    zle redisplay
    return 0
  fi
  cd "$dir"
  local ret=$?
  zle dl-redraw-prompt
  return $ret
}
zle     -N    dl-cd-widget
bindkey '\et' dl-cd-widget

# ALT-R - cd into the selected parent directory
dl-cd-parent-widget() {
  setopt localoptions pipefail 2> /dev/null
  local dir="$(dl -p)"
  if [[ -z "$dir" ]]; then
    zle redisplay
    return 0
  fi
  cd "$dir"
  local ret=$?
  zle dl-redraw-prompt
  return $ret
}
zle     -N    dl-cd-parent-widget
bindkey '\er' dl-cd-parent-widget

fi
