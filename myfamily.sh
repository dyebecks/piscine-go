curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r --argjson 1 '.[$HERO_ID]' '.[] | select(.id==HERO_ID) | .connections | .relatives'