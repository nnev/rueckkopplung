# Rückkopplung

Ask questions during a talk anonymously.


## How it works

Build the software (`go build` should be sufficient) and run it (maybe with `-web.address=myaddress:myport`).
Afterwards, you can go to `http://myaddress:myport` and ask a question there (intended for the audience) and you can go to `http://myaddress:myport/questions` to see all questions asked up to now.

Ideally you set up your system with appropriate surroundings (e.g. DNS and a reverse-proxy to put the `koppler` behind something like `ask.mydomain.com`).


## License

This works is released under the [GNU General Public License v3](https://www.gnu.org/licenses/gpl-3.0.txt). You can find a copy of this license at https://www.gnu.org/licenses/gpl-3.0.txt or in this repository.
