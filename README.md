# lyrical

A CLI for downloading lyrics for any artist

Requires a Spotify App to compile the song list. Lyrics are fetched from a 3rd party API

## Install

Ensure you've create a spotify [app](https://developer.spotify.com/documentation/web-api), they're free and easy to create

Installing the CLI

```
curl -sSL https://raw.githubusercontent.com/mperkins808/lyrical/main/install.sh | sudo bash
```

## Usage

| **Arg** | **Description**                                                                              |
| ------- | -------------------------------------------------------------------------------------------- |
| i       | Your spotify app client id. Optionally set environment variable `SPOTIFY_CLIENT_ID` instead  |
| s       | Your spotify app secret. Optionally set environment variable `SPOTIFY_CLIENT_SECRET` instead |
| d       | directory to store output. eg `./artists`                                                    |
| a       | The artist name to collect songs from. Case Sensitive                                        |

```
lyrical -i <SPOTIFY_APP_CLIENT_ID> -s <SPOTIFY_APP_SECRET> -d ./artists -a Drake
```

## Sample output

Running the above lyrical command would produce an array of the following structure

```json
{
  "artists": ["Drake", "J. Cole"],
  "album": "For All The Dogs Scary Hours Edition",
  "name": "First Person Shooter (feat. J. Cole)",
  "lyrics": "(Pew, pew-pew)\r\nFirst person shooter mode, we turnin' your song to a funeral\r\nTo them niggas that say they wan' off us, you better be talkin' 'bout workin' in cubicles\r\nYeah, them boys had it locked, but I knew the code\r\nLot of niggas debatin' my numeral\r\nNot the three, not the two, I'm the U-N-O\n\nYeah\n\nNumero U-N-O\n\nMe and Drizzy, this shit like the Super Bowl\n\nMan, this shit damn near big as the—\n\n\n\nBig as the what? Big as the what? Big as the what?\n\nBig as the Super Bowl\n\nBut the difference is it's just two guys playin' shit that they did in the studio\n\nNiggas usually send they verses back to me and they be terrible, just like a two-year old\n\nI love a dinner with some fine women\n\nWhen they start debatin' about who the G.O.A.T\n\nI'm like \"Go 'head, say it then, who the G.O.A.T.?\n\n\"Who the G.O.A.T.? Who the G.O.A.T.? Who the G.O.A.T.?\"\n\n\"Who you bitches really rootin' for?\"\n\nLike a kid that act bad from January to November, nigga, it's just you and Cole\n\nBig as the what? Big as the what? Big as the what?(Ayy)\n\nBig as the Super Bowl\n\n\n\nNiggas so thirsty to put me in beef\n\nDissectin' my words and start lookin' too deep\n\nI look at the tweets and start suckin' my teeth\n\nI'm lettin' it rock 'cause I love the mystique\n\nI still wanna get me a song with YB\n\nCan't trust everything that you saw on IG\n\nJust know if I diss you, I'd make sure you know that I hit you like I'm on your caller ID\n\nI'm namin' the album The Fall Off, it's pretty ironic 'cause it ain't no fall off for me\n\nStill in this bitch gettin' bigger, they waitin' on the kid to come drop like a father to be\n\nLove when they argue the hardest MC\n\nIs it K-Dot? Is it Aubrey? Or me?\n\nWe the big three like we started a league, but right now, I feel like Muhammed Ali\n\nHuh, yeah, yeah, huh-huh, yeah, Muhammed Ali\n\nThe one that they call when they shit ain't connectin' no more, feel like I got a job in IT\n\nRhymin' with me is the biggest mistake\n\nThe Spider-Man meme is me lookin' at Drake\n\nIt's like we recruited your homies to beat demon deacons, we got 'em attending a wake\n\nHate how the gang gotta wait for the boss, man, this shit like a prison escape\n\nEverybody steppers, well fuck it, then everybody breakfast and I'm 'bout to clear up my plate (Huh, huh, huh)\n\nWhen I show up, it's motion picture blockbuster\n\nThe G.O.A.T. with the golden pin, the top toucher\n\nThe spot rusher, sprayed his whole shit up, the crop duster\n\nNot Russia, but apply pressure\n\nTo your cranium, Cole's automatic when aimin' 'em\n\nWith The Boy in the status, a stadium\n\nNigga\n\nPart II\n\n\n\nAyy, I'm 'bout to—, I'm bout to—\n\nI'm 'bout to—, yeah\n\nYeah\n\n\n\nI'm 'bout to click out on this shit\n\nI'm 'bout to click, woah\n\nI'm 'bout to click out on this shit\n\nI'm 'bout to click, woah\n\nI'm down to click down you hoes and make a crime scene\n\nI click the trigger on the stick like a high beam\n\nMan, I was Bentley wheel whippin' when I was nineteen\n\nShe call my number, leave her hangin', she got dry-cleaned\n\nShe got a Android, her messages is lime green\n\nI search one name, and end up seein' twenty tings\n\nNadine, Christine, Justine, Kathleen, Charlene, Pauline, Claudine\n\nMan, I pack 'em in this phone like some sardines\n\nAnd they send me naked pictures, it's the small things\n\nYou niggas is still takin' pictures on a dog stream\n\nMy youngers richer than you rappers and they all stream\n\nI really hate that you been sellin' them some false dreams\n\nMan, if your pub was up for sale, I buy the whole thing\n\nWill they ever give me flowers? Well, of course not\n\nThey don't wanna have that talk, 'cause it's a sore spot\n\nThey know The Boy the one they gotta boycott\n\nI told Jim and Jammer I use a GRAMMY as a door stop\n\nGirl gave me some head because I need it\n\nAnd if I fuck with you, then after I might eat it, wait\n\nNiggas talkin' 'bout when this gon' be repeated\n\nWhat the fuck bro? I'm one away from Michael\n\nNigga, beat it, nigga, beat it, what?\n\n\n\nBeat it, what? Beat it, what? Beat it, what? Beat it, what?\n\nBeat it, what? Beat it, what? Beat it, what? Beat it, what?\n\nBeat it, what? Beat it, what? Beat it, what? Beat it, what?\n\nDon't even pay me back on none them favors, I don't need it"
}
```
