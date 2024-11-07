# Übung Design Goals von verteilten Systemen

## Vorbereitung

Wähle eine der folgenden Technologien:

- [CockroachDB](https://www.cockroachlabs.com/): Datenbank
- [RabbitMQ](https://www.rabbitmq.com/): Pub-Sub Messaging
- [Zitadel](https://zitadel.com/): Identity Provider

## Fragen

Beschreibe kurz den Nutzen der Applikation und die wichtigsten Features.

Beantworte anschliessend folgende Fragen.

Es ist wichtiger, dass die Aufgaben gründlich gemacht werden, als dass alle Aufgaben erledigt werden.

Die Antworten sollen in einem kurzen Bericht festgehalten werden und der Klasse vorgestellt werden.

Der Bericht ist per E-Mail an den Dozenten einzureichen. Er wird nicht bewertet, es werden jedoch folgende formellen Anforderungen gestellt:

- Dateiformat: [Markdown](https://www.markdownguide.org/) oder daraus [generiertes PDF](https://pandoc.org/).
- Diagramme: [PlantUML](https://plantuml.com/de/), [Mermaid](https://mermaid.js.org/) o.ä.

Für die Präsentation wird empfohlen, ebenfalls Markdown Dokumente zu verwenden und z.B. mit [Marp](https://marp.app/) zu rendern.

### Design-Ziele

Untersuche die Features der gewählten Technologie und vergleiche diese mit den Design-Zielen von verteilten Systemen (siehe Van Steen Kapitel 1.2).

- Welches sind die wichtigsten Features?
- Für welche von diesen Features werden verteilte Systeme eingesetzt?
- Wie wird dies umgesetzt?
- Welche der folgenden Eigenschaften sind **Transparent**, welche nicht?
  (Zugriff, Standort, Replikation, Änderung des Standorts, gleichzeitiger Zugriff von verschiedenen Usern, Fehler)
- Ist das System **offen**?
- Wie **zuverlässig** ist das System?
- Wie wird die **Sicherheit** gewährleistet?
- In welchem Rahmen kann das System **skalieren**? Wie wird dies umgesetzt?

## Quellen

van Steen 2023
: M. van Steen and A.S. Tanenbaum, Distributed Systems, 4th ed., distributed-systems.net, 2023.