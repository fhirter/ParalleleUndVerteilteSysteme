# Design Goals: Cockroach DB

## Welches sind die wichtigsten Features des untersuchten Systems?

There are many reasons to use CockroachDB, including:

- Resiliency
- Scalability
- Strong consistency
- Geo-partioning and multi-region features
- PostgreSQL-compatibility

## Für welche von diesen Features werden verteilte Systeme eingesetzt?

Resilienz und Skalierbarkeit

## Wie wird dies im untersuchten System umgesetzt?

### Resilienz

Zuverlässigkeit ist ein zentrales Feature von CockroachDB. Die Daten werden mehrfach repliziert, bei einem Ausfall von
einem Node sind die Daten weiterhin verfügbar. Ist ein Node wieder online, werden die Daten automatisch wieder
repliziert.

### Skalierbarkeit

CockroachDB kann mit einer theoretisch beliebigen Anzahl an Nodes betrieben werden, auf welche die Last verteilt wird.
So kann in sehr grossem Ausmass skaliert werden. Möglich ist dies, weil die einzelnen Nodes, anders als bei
traditionelleren Datenbanksystemen, sowohl schreibend als auch lesend arbeiten können.

## Welche der folgenden Eigenschaften sind **Transparent**, welche nicht (vgl. Van Steen Figure 1.3)?

### Zugriff

Der Zugriff ist unabhängig vom zugrundeliegenden Dateisystem oder Betriebssystem.

### Standort

CockroachDB unternimmt einiges, um die verteilte Architektur zu verstecken. Requests werden zum geographisch
naheliegendsten Node geleitet, sowohl schreibend als auch lesend. So bleiben die Latenzen tief, auch wenn die Nodes
global verteilt sind.

### Replikation

Die Replikation der Daten auf den Nodes ist vollständig transparent.

### Änderung des Standorts

Dies ist nur bedingt transparent. Nodes werden fix an einem Standort provisioniert.

### gleichzeitig er Zugriff von verschiedenen Usern

Dies ist vollständig transparent. 

### Fehler

Fehler sind grösstenteils Transparent. 
Bei einem Ausfall von einem Node wird der Service nicht unterbrochen.

## Ist das System **offen** (vgl. Van Steen Kapitel 1.2.3)?

CockroachDB bietet als Schnittstelle den PostgreSQL Dialekt an, welcher sehr verbreitet ist. Dadurch ist das System 
sehr offen und kann vielerorts als drop-in Ersatz für PostgreSQL verwendet werden.

## Wie **zuverlässig** ist das System (vgl. Van Steen Kapitel 1.2.4)?

Siehe "Resilienz"

## In welchem Rahmen kann das System **skalieren**? Wie wird dies umgesetzt (vgl. Van Steen Kapitel 1.2.6)?

Siehe "Skalierbarkeit"