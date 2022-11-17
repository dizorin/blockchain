## *Overview*

This project describes an implementation of blockchain technology for  a program that can register in-game achievements and game related world records of different types such as *“fastest game completion”*, *“best score in a level”* *etc*. This implementation suggestion is meant to provide a *highly trusted achievement register system* that can be used by different game developers, game markets and gamers.


## *Architecture*

- *TODO*


## *Product interaction*

The idea is that the program will need to be integrated into a game with an active internet connection. When the desired game achievement/record gets registered “in-game” and the game has controlled that all requirements to register the achievement were fulfilled the program will be notified, and then it will try to add it in the blockchain.


## *Use cases*

The use cases of the program have to be defined by the game developers. It is also possible to allow the gamers to create/propose their own defined types of achievements/scores, the potential use of this program is almost unlimited.  The list below provides examples of typical “achievements/records”:

* Game bosses
    - Individual boss kill time
    - Individual boss kill score
* Game levels
    - Level clear time
    - Level score
    - Specific items collected
* Game
    - Time to complete the entire game - (For speedrunners community)
    - Amount of received achievements
* Achievement
    - Time of receiving the achievement


## *Security requirements*

- It is seen as a minimum requirement to have an active internet connection to minimize the attempts to register an offline fraudulently received achievement.
- The more games use our program - the more secure our data in the blockchain will be.
- The responsible one for validating the "truthfullness" of the registered achievements/scores is the game, not the program.
- TODO
    - Obviously more requirements will be added after the actual implementation of the program....


## *Target user group characteristics*

The target user group are game developing companies. That is due to the fact that it is them who define what kind of achievements and records the game will have. The developers that will work with the integration of the program into the game will need to have a minimal knowledge of how the blockchain technology works to be able to use the API of this program.


## *Limitations*

At this stage the only limitations are:
- The choise of architecture
    - Will the game developers choose to run the program on their servers or in every game on client-side.
- Where to save the data?
    - Game servers?
    - Separate server?
    - Client-side?