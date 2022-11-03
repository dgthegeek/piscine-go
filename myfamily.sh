curl -s https://learn.zone01dakar.sn/assets/superhero/all.json | jq -r --arg id "$HERO_ID" ' .[] | select( .id=='$id') | .connections.relatives' | sed 's/\"//g'
