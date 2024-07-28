# Todori API

This repo _WILL (Heavily WIP)_ contain an api in Go for future application called Todori.

The aim is to learn Go language and how to write complex api with it.

## About

Todori is a platform where user can plan their work/study plan with something like Kanban boards. They can also collab with other users by inviting them to their workspace.

Current scheme of the v1 api:

![Planning - v1](/assets/schemev1.jpg)

## Getting started

To get started on development, create and fill `.env` (see `.env.example`) and run commands:

```bash
go mod tidy
go generate ./ent # Run "go mod tidy" again if needed
docker-compose up postgres # For local postgres setup
go install github.com/air-verse/air@latest
air
```
