#!/bin/bash

# Read the current version from appversion file
CURRENT_VERSION=$(cat appversion)

# Bump the version
NEW_VERSION=$(echo $CURRENT_VERSION + 0.1 | bc)

# Update the appversion file with the new version
echo $NEW_VERSION > appversion

# Update the values.YAML file with the new version
YAML_FILE="charts/values.yaml"  

# Use sed to replace the version in the values.YAML file
sed -i "s/\(image:\s*tag:\s*\)\"[^\"]*\"/\1\"$NEW_VERSION\"/" $YAML_FILE

# Commit the changes
git add appversion $YAML_FILE
git commit -m "Bump version to $NEW_VERSION"

# Tag the commit
TAG_NAME="v$NEW_VERSION"
git tag $TAG_NAME