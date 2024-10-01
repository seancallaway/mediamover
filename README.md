<div align="center">

  <img src="docs/mediamover_logo.jpeg" alt="logo" width="200" height="auto" />
  <h1>MediaMover</h1>

  <p>A tool for organizing your media files.</p>

  <!-- Badges -->
  <p>
    <a href="https://github.com/seancallaway/mediamover/graphs/contributors">
      <img src="https://img.shields.io/github/contributors/seancallaway/mediamover" alt="contributors" />
    </a>
    <a href="#">
      <img src="https://img.shields.io/github/last-commit/seancallaway/mediamover" alt="last commit" />
    </a>
    <a href="https://github.com/seancallaway/mediamover/issues/">
      <img src="https://img.shields.io/github/issues/seancallaway/mediamover" alt="open issues" />
    </a>
    <a href="https://github.com/seancallaway/mediamover/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/seancallaway/mediamover.svg" alt="license" />
    </a>
  </p>

  <!-- Main Links -->
  <h4>
    <a href="https://github.com/seancallaway/mediamover/issues/">Report Bug</a>
    <span> · </span>
    <a href="https://github.com/seancallaway/mediamover/issues/">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->
# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  * [Screenshots](#camera-screenshots)
  * [Features](#dart-features)
- [Getting Started](#toolbox-getting-started)
  * [Prerequisites](#bangbang-prerequisites)
  * [Installation](#gear-installation)
- [Usage](#eyes-usage)
- [Roadmap](#compass-roadmap)
- [Contributing](#wave-contributing)
  * [Code of Conduct](#scroll-code-of-conduct)
<!-- - [FAQ](#grey_question-faq) -->
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)

<!-- About the Project -->
## :star2: About the Project

<!-- Screenshots -->
### :camera: Screenshots

<div align="center"> 
  <h3>Coming soon!
</div>

<!-- Features -->
### :dart: Features

- Identifies and organizes TV Show files by show and season.
- Identifies and organizes Movie files by genre.

<!-- Getting Started -->
## 	:toolbox: Getting Started

<!-- Prerequisites -->
### :bangbang: Prerequisites

You must have a valid API key for TVDB.

<!-- Installation -->
### :gear: Installation

Download the binary from our [Releases Page](https://github.com/seancallaway/mediamover/releases).

We'll have better installation options in the future.

<!-- Usage -->
## :eyes: Usage

After installing create a `config.ini` file in the following format:

```ini
api_key = <YOUR_TVDB_API_KEY>
tv_root = <The Path to Store Your TV Show Files>
movie_root = <The Path to Store Your Movie Files>
```

After that's created and the values are set properly, you can load your movies or shows.

### :tv: TV Shows

```bash
$ mediamover -v tvshows ~/Downloads/TVShows
Wrote ~/TV/Only Murders in the Building/Season 04/Only Murders in the Building S04E04.mkv
Wrote ~/Videos/TV/Only Murders in the Building/Season 04/Only Murders in the Building S04E05.mkv
Wrote ~/Videos/TV/SEAL Team/Season 07/SEAL Team S07E05.mkv
Unable to write file ~/Videos/TV/The Rookie/Season 06/The Rookie S06E07.mkv: "~/Videos/TV/The Rookie/Season 06/The Rookie S06E07.mkv already exists"
```

### :movie_camera: Movies

_Coming soon._

<!-- Roadmap -->
## :compass: Roadmap

* [ ] User-configurable organization ([#2](https://github.com/seancallaway/mediamover/issues/2))

<!-- Contributing -->
## :wave: Contributing

<a href="https://github.com/seancallaway/mediamover/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=seancallaway/mediamover" />
</a>

Contributions are always welcome!

See [CONTRIBUTING.md](CONTRIBUTING.md) for ways to get started.

<!-- Code of Conduct -->
### :scroll: Code of Conduct

We don't have one yet, but that'll change. For now, be good humans.

<!-- License -->
## :warning: License

Distributed under the Apache 2.0 License. See [LICENSE](LICENSE) for more information.

<!-- Contact -->
## :handshake: Contact

Your Name - [@seancallaway](https://mastodon.social/@seancallaway) - seancallaway@gmail.com

Project Link: [https://github.com/seancallaway/mediamover](https://github.com/seancallaway/mediamover)

<!-- Acknowledgments -->
## :gem: Acknowledgements

**mediamover** wouldn't be possible without the following packages:

 - [regexp2](https://github.com/dlclark/regexp2)
 - [Cobra](https://cobra.dev/)
 - [Viper](https://github.com/spf13/viper)
