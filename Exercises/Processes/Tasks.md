# Übung Prozesse

## Lernziele und Kompetenzen

- Die Studierenden wissen, wie Betriebssysteme die Ausführung von mehreren Prozessen ermöglichen.
- Die Studierenden kennen Kommunikations- und Koordinationstechniken zwischen Prozessen und Systemen.
- Die Studierenden können einfache parallele Applikationen implementieren.
- Die Studierenden können evaluieren, ob für gegebene Anforderungen ein paralleles System notwendig ist.

## Vorgehen

Halte dich an das folgende, [wissenschaftliche Vorgehen](https://en.wikipedia.org/wiki/Scientific_method):

1. **Beobachtung** festhalten und **Fragestellung** formulieren
2. **Recherche**
3. **Hypothese** formulieren
4. **Testen** der Hypothese mit **Experimenten**
5. **Experimente auswerten**
6. **Schlussfolgerungen** und Ausblick formulieren

Stelle sicher, dass du alle Fragen mit **belastbaren Daten** beantworten kannst.

Halte die Erkenntnisse in einem [Markdown-Dokument](https://www.markdownguide.org/) fest.
Schreibe zu jedem der oben genannten Punkte einen Abschnitt. Der Ablauf kann beliebig oft wiederholt werden. 
Entsprechend klein kann der Fokus der Fragestellung sein.

**Schicke das Dokument nach dem Unterricht an den Dozenten.**

## Aufgaben

### Systemanalyse

Verbinde dich mit den separat kommunizierten Zugangsdaten mit dem Linux-System.

Untersuche anhand der laufenden Prozesse, wozu das System dient.
Identifiziere die wichtigsten Prozesse.

Untersuche diese identifizierten Prozesse mit `ps` und `htop` indem du folgende Fragen beantwortest.
Beschränke dich dabei nicht nur auf die reine Beantwortung der Frage, sondern erkläre die Frage an sich sowie deinen
Lösungsprozess.

- Welche Prozesse nutzen wie viel Speicher?
- Welche Prozesse nutzen wie viel CPU?
- Welche Prozesse haben wie viel CPU Zeit verwendet?
- Welchen Ausführungsstatus haben die Prozesse?
- Welche Scheduling Priorität haben die Prozesse?
- Wie viele Prozesse starten die einzelnen Applikationen?
- Welche Prozesse nutzen mehrere Threads?

### Multiprocessing

Entwickle ein Programm, das den Mittelwert einer Liste von Zahlen berechnet.

Das Programm soll eine Liste von zufälligen Zahlen generieren und davon den Mittelwert berechnen.

Verwende für die Lösung [Go](https://go.dev/), dieses bietet sehr einfaches Multiprocessing.

Nutze die [beiliegende Vorlage](Multiprocessing.go) für die Berechnungen.

Beantworte gemäss oben beschriebenem Vorgehen folgende Fragen:

- Welcher Zusammenhang besteht zwischen der Berechnungszeit und der Länge der Zahlenreihe sowie der gewählten 
  Nebenläufigkeit (Anzahl nebenläufige Berechnungen)?
- Welche Nebenläufigkeit ist für welche Länge der Zahlenreihe optimal?

Recherchiere zuerst, wie Go Channels funktionieren und vergleiche mit der Beschreibung im Kapitel 3.6 (Silberschatz, 
2019). Erläutere den Code der Vorlage.

### Scheduling

Löse folgende Aufgabe für verschiedene Scheduling Algorithmen:

- [Priority based, preemptive, round robin](https://en.wikipedia.org/wiki/Fixed-priority_pre-emptive_scheduling)
- [Completely Fair Scheduler](https://en.wikipedia.org/wiki/Completely_Fair_Scheduler)
- [Multilevel Feedback Queue](https://de.wikipedia.org/wiki/Multilevel_Feedback_Queue)

| Bezeichner | Priorität | CPU Zeit | Ready Zeit |
|------------|-----------|----------|------------|
| P1         | 40        | 20       | 0          |
| P2         | 30        | 25       | 25         |
| P3         | 30        | 25       | 30         |
| P4         | 35        | 15       | 60         |
| P5         | 5         | 10       | 100        |
| P6         | 10        | 10       | 105        |

Jeder Prozess hat eine Priorität, wobei eine höhere Zahl eine höhere Priorität bedeutet.
Zusätzlich zu den aufgeführten Prozessen existiert ein Idle-Task, der ausgeführt wird, wenn keine anderen Prozesse
verfügbar sind.
Er besitzt Priorität 0.

Der Scheduler wird alle 10 Zeiteinheiten ausgeführt.

1. Zeichne ein Gantt Diagramm für die Prozessabfolge.
2. Wie hoch ist die CPU-Auslastung?
3. Wie gross ist die Wartezeit für jeden Prozess?
