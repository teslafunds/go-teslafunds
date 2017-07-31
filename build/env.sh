#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
<<<<<<< HEAD
ethdir="$workspace/src/github.com/teslafunds"
if [ ! -L "$ethdir/go-teslafunds" ]; then
=======
ethdir="$workspace/src/github.com/dubaicoin-dbix"
if [ ! -L "$ethdir/go-dubaicoin" ]; then
>>>>>>> 7fdd714... gdbix-update v1.5.0
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../../../. go-teslafunds
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$ethdir/go-teslafunds"
PWD="$ethdir/go-teslafunds"

# Launch the arguments with the configured environment.
exec "$@"
