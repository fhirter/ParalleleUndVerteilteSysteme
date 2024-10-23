# Übung Prozesse

## Lernziele

- Die Studierenden wissen, wie Betriebssysteme die Ausführung von mehreren Prozessen ermöglichen.

## Aufgabe

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
