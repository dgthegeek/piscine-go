curl -s https://learn.zone01dakar.sn/assets/superhero/all.json | jq -r --arg HERO_ID "$HERO_ID" ' .[] | select(.id=='$HERO_ID') | .connections.relatives' | tr -d '\n'
