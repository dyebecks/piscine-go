export HERO_ID=1
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r '.[] | select(.id==1) | .connections | .relatives'
