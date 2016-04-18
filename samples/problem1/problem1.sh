#!/bin/bash

Usage () {
  echo "Usage:"
  echo "./bash_example [--help|-h]"
  echo "./bash_example --create-file=<filename> [--no-prompt] [--verbose]"
}

Info () {
  Usage
  echo "--no-promt : ........"
  echo "--verbose : ....... "
}

Inquire () {
  echo  -n "$1 [y/n]? "
  read answer
  finish="-1"
  while [ "$finish" = '-1'  ]
  do
    finish="1"
    if [ "$answer" = ''  ];
    then
      answer=""
    else
      case $answer in
        y | Y | yes | YES ) answer="y";;
        n | N | no | NO ) answer="n";;
        *) finish="-1";
          echo -n 'Invalid response -- please reenter:';
          read answer;;
      esac
    fi
  done

  [ "$answer" = "n" ] && return "1"
  [ "$answer" = "y" ] && return "0"
}

Overwrite() {
  rm $1
  touch $1
}

CreateFile() {
  if [ -f $1  ]; then
    echo "File already exists!"
    Inquire
    answer=$?
    [ "$answer" -eq "0" ] && Overwrite $1
    exit 1
  fi

  touch $1
}

if [ -z "$1" ]; then
  echo "Empty parameters" && Usage
  exit
fi

for arg in "$@"
do
  case "$arg" in
    --create-filename)
      echo 'Filename not specified'
      shift
      ;;
    --create-filename=*)
      CREATEFILENAME="${arg#*=}"
      shift
      ;;
    -h|--help)
      Info && exit
      ;;
    --no-prompt)
      NOPROMPT=true
      shift
      ;;
    --verbose)
      VERBOSE=true
      shift
      ;;
    *)
      Usage
      exit
      ;;
  esac
done

if [ -z "${CREATEFILENAME}" ]; then
  Usage && exit
fi

echo "FILE EXTENSION  = ${CREATEFILENAME}"
echo "NOPROMPT  = ${NOPROMPT}"
echo "VERBOSE  = ${VERBOSE}"

CreateFile $CREATEFILENAME

echo 'End of the line'
exit 0
