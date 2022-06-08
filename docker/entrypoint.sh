#!/bin/sh
set -e
if [[ ! -f /opt/workflow/daisy/config/settings.yml ]]
then
    cp /opt/workflow/daisy/default_config/* /opt/workflow/daisy/config/
fi
if [[ -f /opt/workflow/daisy/config/needinit ]]
then
    /opt/workflow/daisy/daisy init -c=/opt/workflow/daisy/config/settings.yml
    rm -f /opt/workflow/daisy/config/needinit
fi
/opt/workflow/daisy/daisy server -c=/opt/workflow/daisy/config/settings.yml
