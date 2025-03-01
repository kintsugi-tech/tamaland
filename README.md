# Tamaland

## What is it?

Tamaland is an OnChain version of the classic Tamagotchi game. The goal of the game, as in the classic version, is to keep your tamagotchi alive for as long as possible. As time passes, the player will see their tamagotchi level up. Will they be skilled and patient enough to reach the highest levels?

This version of the game represents the first MVP of the project.
At Kintsugi, we have thought of many ideas to make the game even more fun and immersive, and these ideas will be presented at the end of this documentation.
We highlight here the main one in this introductory phase: each tamagotchi can become a living and defined entity, represented by an NFT that can be exchanged between players, who will take turns caring for it. Each player change, like a world switch, will bring unique evolutionary differences to the tamagotchi!

### Technology stack

The technology stack that makes up the game is based on three main components: WorldEngine's Game Shard [Cardinal](https://world.dev/cardinal/introduction) (a sharded roll-up SDK: game-shard framework used for building the game's backend), [Unity](https://unity.com/) (the game's client), and [Nakama](https://heroiclabs.com/docs/nakama/getting-started/index.html) (game server to facilitate communication between Cardinal and Unity).

#### World Engine (by [ArgusLabs](https://argus.gg/))

World Engine its the project's core. We used here its Cardinal framework:

> "Cardinal is a World Engine game shard framework for building powerful, highly scalable onchain games.
> Through the World Engine’s shard communication interface (SCI), Cardinal interoperates seamlessly with EVM smart contracts on the World Engine stack, providing web2 scale performance for onchain games while preserving interoperability."
>
> [World Engine docs](https://world.dev/cardinal/introduction)

World Engine represents a truly innovative development framework in the field of Web3 game development.
The strength of Cardinal is its ability to handle thousands of game updates per second, creating a smooth real-time experience by producing a high number of ticks (~20/30) per second. 

Since each tick represents a change in the world's state, this translates, in blockchain terms, to Cardinal's capability of generating up to 20/30 blocks per second—an enormous leap forward for on-chain game development.

##### Cadinal Editor

Cardinal is based on an ECS (Entity Component System) architecture. Entities, components, and systems can be easily created within the game and modified through a development editor, the Cardinal Editor.

The Cardinal Editor has been extensively used in nearly every phase of development due to its completeness and convenience.

#### Nakama (by [Heroic Labs](https://github.com/heroiclabs/nakama))

Nakama is an open-source game server framework. It has been used here to facilitate the connection between the game client and Cardinal.

Its adoption was driven by the recommendation of ArgusLabs, which suggests using Nakama due to its available plugins for various game engines.

Through Nakama, it was possible to establish a stable connection between the game client and the Cardinal shard, ensuring a seamless experience for players so they can fully enjoy the game.

#### Unity (by [Unity Technologies](https://unity.com/))

Unity is the highly popular game engine developed by Unity Technologies, offering the ability to create both 2D and 3D worlds.

In this project, it was used to create a small corner of the world in 2D, featuring a charming retro-style aesthetic.

## How does it work?

### Tamagotchi creation

Each Tamagotchi is defined by a set of parameters that determine its in-game statistics. The parameters included in the game are:

- Health (Hp): Represents the Tamagotchi's life.
- Energy (E): Indicates the character's energy.
- Food (Fd): Represents the Tamagotchi’s satiety, determining how often it needs to eat.
- Level (L): Indicates the character’s level.
- Status (S): Represents the character’s current state.

Hp, E, and Fd range from a minimum of 0 to a maximum of 100. L starts at a minimum of 1 and increases over time, while the initial Status is set to Normal.

Each Tamagotchi has a display_name, which represents its name, and a nickname, which serves as a unique identifier for the character within the game.

The nickname is set as the UserId assigned to the player upon their first connection to Nakama. This is crucial for linking the user to their character every time they reconnect.

However, to create a Tamagotchi, a player must also have a "Persona" that identifies them within Cardinal. This Persona will contain the address that allows the player to perform and sign transactions on the blockchain.

The key messages required for Tamagotchi creation are two:

- MsgClaimPersona: Allows the creation of a Persona linked to the player's account. The required payload is:

```json
{"personaTag": "{session.UserId}"}
```
For this MVP, it was decided to automate the Persona name by assigning it the user's UserId.

- MsgCreatePlayer: Allows the creation of a new Tamagotchi. The required payload is:
```json
{"personaTag": "{session.UserId}", "nickname": "{session.UserId}", "display_name": "New-Player"}
```
personaTag specifies which address will sign the transaction.
Nickname and display_name are the required fields for the new character.
Hp, E, and Fd are automatically initialized to their maximum values, while L starts at 1.

### Game Flow

The game architecture in Unity consists of three scenes:

- Welcome Scene: This is the startup screen displayed when a player connects to the game. It contains two buttons:
  - New Game: Starts the game and creates a new Tamagotchi if the player hasn't created one yet. If a Tamagotchi already exists, it reloads the statistics of the previous character.
  - Load Scene: Currently disabled. The underlying idea is to separate the logic within New Game, allowing the player to choose whether to load an existing Tamagotchi or create a new one.
- Main Scene: This is the core gameplay screen, where the player's Tamagotchi is displayed alongside status bars for health, hunger, and energy. In the bottom right corner, the character’s level is shown, which, as mentioned earlier, depends on the character's lifetime. Additionally, there are two buttons:
  - Eat: Sends a Tx to Cardinal to trigger the "Eat" action.
  - Sleep: Sends a Tx to Cardinal to trigger the "Sleep" action.</br>
  These two buttons are essential as they restore the hunger and energy of the character, affecting the mechanics described in the following section.
- Dead Scene: This scene is displayed when the character’s status is set to "Dead". It features a New Game button that allows the player to overwrite their old character and restart the game.

### Cardinal loop system

The aging process of the Tamagotchi is managed through Cardinal's systems.

Each Tamagotchi has an internal "LastUpdate" parameter, which stores the last time the Tamagotchi was updated.

A global game-level parameter determines how often Cardinal should update all Tamagotchis. At each tick, Cardinal checks whether it is time to update the character’s statistics.

Using this logic: If a Tamagotchi does not need an update during a specific tick, the system waits for the next one.
This process continues until all Tamagotchis have been updated.
During each tick, the following systems are executed.

#### Energy System

The EnergySystem decreases a portion of the Tamagotchi's energy every `StatUpdateInterval` seconds. The lost energy can be restored through the "Sleep" action. If the character is asleep, the EnergySystem will add energy instead of subtracting it.

#### Food System

The FoodSystem decreases a portion of the Tamagotchi's food every `StatUpdateInterval` seconds. The lost food can be restored through the "Eat" action. If the character is eating, the FoodSystem will add food instead of subtracting it.

#### Health System

The HealthSystem decreases a portion of the Tamagotchi's health if its energy or food levels drop below a certain threshold (80%). The hungrier or more exhausted the Tamagotchi is, the greater the health loss will be.

When health reaches zero, the Tamagotchi will unfortunately die.

#### Status System

The StatusSystem manages the state transitions of the Tamagotchi. There are four possible character states:

- *Normal:* This is the default state of the character. In this state, the Tamagotchi can perform actions such as eating or sleeping.
- *Eating:* This state indicates that the Tamagotchi is currently eating. While in this state, no other actions can be performed.
- *Sleeping:* This state indicates that the Tamagotchi is sleeping. While in this state, no other actions can be performed.
- *Dead:* This state indicates that the Tamagotchi has unfortunately died, and it will no longer be possible to play with the character.
Each status is associated with an "EndStateTimestamp," which marks the moment when the Tamagotchi will exit that state.
Except for death, which is a permanent state, the StatusSystem will check at every tick whether the Tamagotchi is ready to return to the "Normal" state.

### Level System

The LevelSystem manages the Tamagotchi's level progression. It uses a different time interval from `StatUpdateInterval` and is based on a different LastUpdate value (instead of LastUpdate.Timestamp, it uses LastUpdate.LevelUpdateTimestamp). This ensures that level progression is slower compared to the Tamagotchi's stat updates.

### LastUpdate System

The LastUpdateSystem is the final system called into action, managing the update of "LastUpdate."
When the `StatUpdateInterval` is reached, the Tamagotchi's LastUpdate field is also updated through this system.

The LastUpdateSystem does not update LastUpdate.LevelUpdateTimestamp: this is handled by the LevelSystem.

### Cardinal Msgs

When a Tamagotchi can perform actions, they are communicated to Cardinal by the client through "Eat" or "Sleep" messages. These messages only contain the identifier of the Tamagotchi they refer to (the nickname).

The message will only trigger an instantaneous state change in the Tamagotchi, while the respective systems will manage the recovery of the associated stat.

When the player wants to start a new game (e.g., after the Tamagotchi dies), they can use the game client to send a "Respawn" message, which will delete the old Tamagotchi and create a new one.

### Client<>Backend Integration

Since everything is managed by Cardinal, the game client queries Cardinal at every tick to retrieve the player's Tamagotchi's current stats and displays them in real-time to the player.
When the player wants to recover food or energy, they can use the respective UI buttons to send the corresponding messages to Cardinal.

### How to play?

To start playing, it is essential to follow the [Quickstart Guide](https://world.dev/quickstart) by Argus Labs regarding the installation of Go, World CLI, and Docker or Orbstack (recommended).
Once these steps are completed, simply run the command: </br>
``` world cardinal start ``` </br>
which will start both Cardinal and Nakama. Once they are running, all you need to do is open the game executable and have fun!

## Future versions

In future versions, we would like to introduce the possibility of turning the Tamagotchi into a tradable NFT between players.
Each player will have a different type of world (mountainous, marine, plain, etc.), and each Tamagotchi will be a different little creature.
Trading Tamagotchis between different worlds will lead to initially unpredictable mutations in their appearance and behavior, making the experience even more enjoyable!
Furthermore, Tamagotchis that, unfortunately, have died and become ghosts will be usable in other themed games!