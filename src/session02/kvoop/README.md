## Übung: Key-Value Store (Cont.)

* Schreibe einen kleines Programm `kv`, mit dem Du Schlüssel-Werte Paare in einer Datei speichern und abfragen kannst.
* Baue Deinen KV-Store so (um), dass er intern eine Klasse Store verwendet, die die Operationen auf den internen Storage abstrahiert.
* Teste die Klasse Store.

Installieren:
```shell
cd /path/to/go-workshow
go install session01/kv
``` 

Setzen von Werten:
```shell
kv name=Mancke vorname=Sebastian alter=42
```

Abfragen bestimmter Werte:
```shell
kv name vorname
> name=Mancke
> vorname=Sebastian
```

Abfragen aller Werte:
```shell
kv
> name=Mancke
> vorname=Sebastian
> alter=42
```