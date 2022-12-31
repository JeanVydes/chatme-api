# FEATURES OR IDEAS

## Money

- The main idea is get a fee per suscription (2%) or per message (5%)
- No ADS

## Simples but Extensive

- Authentication
- OAuth2 support for Google and Discord
- Profiles
- Support DM and Groups
- Set up of chat adquision method (suscription per month, price per messages) (prices are set by users)
- Feed
- Payment Wall
- Support Ways
- Refund suscription or bulk message (or 1 messages if every one has a price more than 5.00 USA Dollars)

## Well... I dont lost anything trying

- Machine Learning model to predict more likeable people based on followed people and bought suscription

- Mobile App (in a very long future, i dont know Java)

## Very Specific

- Authentication need to be super secured, store password as sha256, implementation OAuth like SMS verification codes, emails codes, OAuth2 Apps codes; 

- SQL for basic static not frequently modified data as email, username, account creation data, and like that; use NoSQL (MongoDB) for frequently modified data and quickly access as preferences, suscriptions, activity.

- Chats are a hard topic, I could suggest use Redis for a faster access to messages from API, but we need persistent data to long-term, so is need store data into NoSQL (performance and uses decision i think), so think later a way to implement this.