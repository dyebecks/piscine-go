curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg a=="HERO_ID" '.[] | select(.id==$a) | .connections | .relatives'