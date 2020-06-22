![Build](https://github.com/scoir/canis/workflows/Build/badge.svg)

# CANIS

Canis is an extensible credentialing platform providing a configurable, easy to use environment for issuing and verifying credentials that conform with decentralized 
identity standards including [W3C decentralized identifiers](https://w3c.github.io/did-core/) (DIDs), [W3C DID resolution](https://w3c-ccg.github.io/did-resolution/), and [W3C verifiable credentials](https://w3c.github.io/vc-data-model/).

## Summary

- [**Why Canis?**](#why-canis)
- [**Features**](#features)
- [**Distributions**](#distributions)
- [**Development**](#development)
- [**License**](#license)

## Why Canis?

Issuing digital credentials requires an institution or organization to have an agent to represent its interest in the digital landscape.
This agent must act as a fiduciary on behalf of the organization, must hold cryptographic keys representing its delegated authority, and it must communicate
via [DIDComm Protocols](https://github.com/hyperledger/indy-hipe/pull/69).  

In addition needing an agent, organizations need a way to instruct this agent how to issue a credential and to whom.  That requires information that is currently stored 
in legacy (in ToIP terms) systems.

Canis serves as a platform for creating, launching and empowering agents to participate in a credentialing ecosystem on any organization's behalf.  In addition,
Canis provides an easy to use RESTful API and extensible data model to allow for endorsing agents on behalf of any hierarchy of organizational structure.

## Features
1. **REST API**: Canis can be operated with its RESTful API for maximum flexibility
1. **Multiple Databases**: Canis can be used with Mongo or Postgres
1. **Multiple DID Resolution**: DID resolution can be performed against...
1. **Multiple VC Formats**: Issue, prove and verify CL, JWT and JSON-LD credentials, even in the same issuance
1. **Multiple Ledger Support**:  Credential issuing on Indy, ?
1. **Plugins**: Extendable architecture for adding functionality to Canis and APIs. 
1. **CLI**: Control your Canis platform from the command line.
1. **Mediator/Router**:  Offer single endpoint to all entities maintaining identity on platform
1. **Mailbox**: Message routing and storage for agents in support of remote, not-always-on devices
