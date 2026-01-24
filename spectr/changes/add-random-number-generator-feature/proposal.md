# Add Random Number Generator Feature

## Overview
Implement IRandomNumberGeneratorFeature from Open-XML-SDK for dependency-injectable random number generation enabling deterministic testing and custom RNG strategies.

## Motivation
Open-XML-SDK provides IRandomNumberGeneratorFeature as part of Features system. Supports testing, compliance with RNG requirements, and reproducible document generation.

## Goals
- Implement IRandomNumberGeneratorFeature interface
- Provide default implementation using crypto/rand
- Support custom RNG implementations via features
- Use feature-based RNG for all internal random operations (paragraph ID generation, relationship ID generation, salt generation for encryption)
- Enable deterministic testing with seeded RNG

## Dependencies
- Depends on: openxml features system
- Related: add-paragraph-id-features

## Estimated Effort
1 week

## Priority
P2
