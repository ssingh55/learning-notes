### Quick guide with a script helps to clean up old snap versions and free some disk space in your Ubuntu systems

# Following are the commands to remove old snap package:

## 1. List installed Snap versions:
   ```
   snap list
   ```

## 2. Remove the desired Snap package (replace "old_snap_package" with the actual package name):
   ```
   sudo snap remove old_snap_package
   ```

## 3. Optionally, clean up residual data:
   ```
   sudo rm -rf /var/cache/snapd/old_snap_package
   ```

## Make sure to replace "old_snap_package" with the name of the Snap package you want to remove. If you have multiple versions installed, repeat the steps for each one.

# Following are the commands to remove old snap package specific revision

## To remove a specific revision of a Snap package in Ubuntu, you can use the `--revision` option with the `snap remove` command. Here's an example:

```bash
sudo snap remove --revision=<revision_number> old_snap_package
```

## Replace "old_snap_package" with the actual name of the Snap package and `<revision_number>` with the revision number you want to remove.

### For instance, if you want to remove revision 3 of a Snap called "firefox," the command would look like this:

```bash
shubham@shubham-prlappy:~$ sudo snap remove firefox --revision=3626
firefox (revision 3626) removed
```
### If space is not getting cleared up then remove the cached directory data as well but be cautious while executing this
# Clean up Snap cache
sudo rm -rf /var/cache/snapd/

## Remember to replace "example_snap" with the name of your Snap package and adjust the revision number accordingly.

# Script to remove all the disabled snap old package

## 1. Open a text editor and create a new file, for example, `remove_disabled_snaps.sh`.

## 2. Add the following lines to the script:

```bash
#!/bin/bash

# Get a list of disabled snaps
disabled_snaps=$(snap list --all | grep disabled | awk '{print $1}')

# Loop through each disabled snap and remove it
for snap in $disabled_snaps
do
    sudo snap remove $snap
done

# Clean up Snap cache
sudo rm -rf /var/cache/snapd/
```

## 3. Save and close the file.

## 4. Make the script executable by running the following command in the terminal:

```bash
chmod +x remove_disabled_snaps.sh
```

## 5. Run the script:

```bash
./remove_disabled_snaps.sh
```

## This script will identify all disabled Snap packages and remove them one by one. Please be cautious and ensure that you want to remove the disabled snaps before executing the script.


# Script to remove all the disabled snap old package revision

```bash
#!/bin/bash

# Get a list of disabled snaps
disabled_snaps=$(snap list --all | grep disabled | awk '{print $1}')

# Loop through each disabled snap
for snap in $disabled_snaps
do
    # Get the list of revisions for the disabled snap
    revisions=$(snap list $snap --all | grep disabled | awk '{print $3}')

    # Loop through each revision and remove the snap
    for revision in $revisions
    do
        sudo snap remove --revision=$revision $snap
    done
done

# Clean up Snap cache be cautious in executing this
sudo rm -rf /var/cache/snapd/
```

### Here's how to use the script:

## 1. Open a text editor and create a new file, for example, `remove_disabled_snap_revisions.sh`.

## 2. Copy and paste the above script into the file.

## 3. Save and close the file.

## 4. Make the script executable by running the following command in the terminal:

```bash
chmod +x remove_disabled_snap_revisions.sh
```

## 5. Run the script:

```bash
./remove_disabled_snap_revisions.sh
```

### This script will identify all disabled Snap packages, retrieve their revisions, and remove each revision one by one. Please use this script carefully, and ensure that you want to remove the disabled snaps along with their revisions before executing it.
