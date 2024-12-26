# ECS - Component Based Entity System

## Entity (Store at most one Component of each type)
    Any Object in the game like 
- Player
- Enemy
- Bullet
- Item

## Component (Pure Data components)
    Any property of an Entity
- Position
- Velocity
- Health
- Damage
- Name
- Sprite

## System
    Any logic that operates on Entities
- Movement
- Collision
- AI
- Render
- Animation
- Input

Example:
Entity: Player
    Components: Position, Velocity, Health, Damage, Name, Sprite, Input
    System: Movement
Entity: Enemy
    Components: Position, Velocity, Health, Damage, Name, Sprite, AI
Entity: Bullet
    Components: Position, Velocity, Damage, Name, Sprite
Entity: Tile
    Components: Position, Sprite


## GameEngine architecture
    The Game Engine is the main class that runs the game.
- Scene
    - System
    - EntityManager
        - Entity
            - Component 
