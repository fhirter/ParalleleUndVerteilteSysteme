#!/bin/bash
export RESTIC_REPOSITORY="/root/backup"
export RESTIC_PASSWORD="yourpassword"

# Directory to backup. Replace with the path you want to back up.
BACKUP_PATH="/root/data"

# Perform the backup
restic backup "$BACKUP_PATH"

# Optional: Clean up old snapshots based on a policy (e.g., keep the last 7 snapshots)
restic forget --keep-last 7 --prune

# Check for repository errors
restic check

# Optional: List all snapshots after backup
# restic snapshots