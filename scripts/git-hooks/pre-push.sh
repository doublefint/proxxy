#!/bin/sh

# An example hook script to verify what is about to be pushed.  Called by "git
# push" after it has checked the remote status, but before anything has been
# pushed.  If this script exits with a non-zero status nothing will be pushed.
#
# This hook is called with the following parameters:
#
# $1 -- Name of the remote to which the push is being done
# $2 -- URL to which the push is being done
#
# If pushing without using a named remote those arguments will be equal.
#
# Information about the commits which are being pushed is supplied as lines to
# the standard input in the form:
#
#   <local ref> <local sha1> <remote ref> <remote sha1>
#

remote="$1"
url="$2"

# Pre-push configuration
remote=$1
url=$2
echo >&2 "Try pushing $2 to $1"

# Run test and stop on first failed
CMD="make lint"

# Run test and return if failed
printf "Running golint..."
$CMD
RESULT=$?
if [ $RESULT -ne 0 ]; then
  echo >&2 "$RESULT"
  echo >&2 "FAILED $CMD"
  exit 1
fi

exit 0