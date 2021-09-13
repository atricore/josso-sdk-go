#!/bin/bash


Title() {
    printf -v Bar '%*s' $((${#1} + 4)) ' ' 
    printf '%s\n| %s |\n%s\n' "${Bar// /-}" "$1" "${Bar// /-}"
}

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

ENV="/opt/atricore/tools/env.sh"
if [ ! -f $ENV ] ; then
    echo "File not found : $ENV"
    exit 1
fi

source $ENV

# This is the pipeline folder name for the JOSSO EE server (i.e. josso-ee-2-m)
if [ -z "$1" ] || [ -z "$2" ] ; then
    echo "Usage: $0 <josso-version> <wb-version>"
    exit 1
fi

PIPELINE=$1
JOSSO_VERSION=$2

BUILD_FOLDER=$SCRIPT_DIR/../../$PIPELINE
if [ ! -d "$BUILD_FOLDER" ] ; then
        echo "Build folder not found : $BUILD_FOLDER"
        exit 1
fi
if [ ! -f "$BUILD_FOLDER/buildNumber.properties" ] ; then
        echo "Build number not found : $BUILD_FOLDER/buildNumber.properties"
        exit 1
fi

BUILD_DATE=`sed 's/^#\(.*\)/\1/;2q;d' "$BUILD_FOLDER"/buildNumber.properties`
BUILD_NUM=`sed 's/^buildNumber=\(.*\)/\1/;3q;d' "$BUILD_FOLDER"/buildNumber.properties`
BUILD_TIMESTAMP=$(date -d "$BUILD_DATE" +%Y%m%d.%k%M%S)
BUILD_TAG=$PIPELINE-$BUILD_TIMESTAMP-$BUILD_NUM
MVN_TARGET="$BUILD_FOLDER"/distributions/josso-ee/target
TMP=/tmp/josso

LOG_FILE=$SCRIPT_DIR/test-$BUILD_TAG.log

JOSSO_UNIX_DIR="$MVN_TARGET"/josso-ee-"$JOSSO_VERSION"-SNAPSHOT-server-unix/josso-ee-"$JOSSO_VERSION"-SNAPSHOT
if [ ! -d "$JOSSO_UNIX_DIR" ] ; then
        echo "JOSSO unix dir not found : $JOSSO_UNIX_DIR"
        exit 1
fi

JOSSO_SERVER=/tmp/josso/josso-ee-"$JOSSO_VERSION"-SNAPSHOT/server

Title "Installing JOSSO server : $JOSSO_SERVER"
mkdir -p /tmp/josso
cp -r $JOSSO_UNIX_DIR /tmp/josso/

if ! cp $SCRIPT_DIR/com.atricore.idbus.console.appliance.default.idau.cfg $JOSSO_SERVER/etc ; then
    echo "ERROR copying appliance default config"
    exit 1
fi

cd $JOSSO_SERVER/bin
./stop

if ! ./start ; then
    echo "ERROR : error starting JOSSO"
    exit 1
fi

Title 'Waiting for JOSSO server ... '

$STARTED_STR="Installable Unit not found in updates"

timeout 180 grep -q "$STARTED_STR" <(tail -f $JOSSO_SERVER/data/log/atricore.log)
grep $STARTED_STR $JOSSO_SERVER/data/log/atricore.log
if [ $? !=  0] ; then
    echo "Server not started!"
    exit 1
fi

