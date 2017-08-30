# go-nagios-config
A Nagios config parser in Go (work in progress).

- The package is currently able to parse a valid nagios configuration and to print it;
- timeperiods are not parsed;

### Roadmap:
- parse timeperiods;
- be able to actually resolve group members and templating hierarchies;
- be able to validate configuration;
- be able to write down configuration in a different format (for example: expand services defined on multiple hosts to single-host service definitions or - on the opposte - collapse multiple single-host definitions of a single service into a single, multi-host, definition);
- suggestions?
