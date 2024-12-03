# Übung Edge Computing

## Lernziele

Die Studierenden

- kennen die Beweggründe und Herausforderungen von parallelen und verteilten Systemen
- Kennen verschiedene dezentrale Architekturstile und deren Anwendung
- eine einfaches verteiltes System entwerfen

## Dokumentation

Halte die Erkenntnisse und Designs in einem [Markdown-Dokument](https://www.markdownguide.org/) mit [Mermaid](https://mermaid.js.org/) fest.

## Aufgabenstellung

Eine Applikation publiziert bei jeder Änderung einen Event (siehe [events.csv](events.csv), der via REST API in 
einem Backend persistiert wird und zur Synchronisation via Websockets an die anderen Instanzen gepusht wird.

Da sich die Instanzen in der Regel im gleichen lokalen Netz befinden, soll untersucht werden, unter welchen Bedingungen ein Edge Server von Vorteil ist.

Der Edge Server hat folgende Aufgaben:

- Verteilen der Events an die lokalen Instanzen
- Batch basiertes persistieren der Events im Backend

Es gelten folgende Rahmenbedingungen:

- Request Latenz zum Backend (Netzwerk und Verarbeitung): 500ms
- Maximale Batch Latenz: 2s

- Filesize: 1284kB
- Event Count: 2826
- 454B per Event

Aufgaben:

- Simuliere den Status Quo (Cloud Server) und bestimme mit Messungen die Performance

Fragen:

- Wie kann der Nutzen des Edge-Servers quantifiziert werden?

## Aufgaben

- 