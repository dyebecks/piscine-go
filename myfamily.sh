echo HERO_ID
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r '.[] | select(.id==1) | .connections | .relatives'