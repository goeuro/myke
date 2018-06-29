#!/usr/bin/env bash

function _myke_task_complete() {
  COMPREPLY=();
  local word="${COMP_WORDS[COMP_CWORD]}";
  COMPREPLY=($(compgen -W "$(myke --tasks | awk -v FS='\t' '{print $1}' | grep -e "^$word" -e "/$word")"));
}

complete -F _myke_task_complete myke
