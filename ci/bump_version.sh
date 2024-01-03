#!/bin/bash

# Read the current version from appversion file
CURRENT_VERSION=$(cat appversion)

# Bump the version
NEW_VERSION=$(echo $CURRENT_VERSION + 0.1 | bc)

# Update the appversion file with the new version
echo $NEW_VERSION > appversion

# Update the YAML file with the new version
YAML_FILE="charts/values.yaml"  

# Use sed to replace the version in the YAML file
sed -i "s/\(tag:\s*\)\"[^\"]*\"/\1\"$NEW_VERSION\"/" $YAML_FILE

# Commit the changes
git add appversion $YAML_FILE

# Commit the changes and tag
git add appversion $YAML_FILE
git commit -m "Bump version to $NEW_VERSION"
git tag -a "v$NEW_VERSION" -m "Version $NEW_VERSION"

# Push changes and tags to the remote repository
git push origin main --tags