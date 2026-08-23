import 'dart:typed_data';

enum Inhabit { foreigner, inhabitant, hermit, forgotten }

class Flavor {
  String raw = '';
  String filtered = '';

  String name = '';
  String race = '';
  String job = '';

  Uint8List? image;

  int day = 0;

  bool get hasFiltered => filtered != '';

  Flavor clone() {
    final flavor = Flavor();
    flavor.raw = raw;
    flavor.filtered = filtered;
    flavor.name = name;
    flavor.race = race;
    flavor.job = job;
    flavor.image = image;
    flavor.day = day;
    return flavor;
  }

  String formatText() {
    final buf = StringBuffer();
    buf.write('名前: $name\n');
    buf.write('種族: $race\n');
    buf.write('職業: $job\n');
    buf.write('特徴の説明: $filtered\n');
    return buf.toString();
  }
}

class PlayerAction {
  String raw = '';
  String filtered = '';

  Uint8List? image;

  int day = 0;

  bool get hasFiltered => filtered != '';
}

class Player {
  Inhabit inhabit = Inhabit.foreigner;
  int settled = 0;
  String id = '';

  List<Flavor> flavors = [];
  Flavor flavor = Flavor();

  List<PlayerAction> actions = [];
  PlayerAction action = PlayerAction();
  bool committed = false;

  String summary = '';

  bool get isForeigner => inhabit == Inhabit.foreigner;

  Flavor get lastFlavor => flavors.last;
  PlayerAction? get lastAction => actions.isEmpty ? null : actions.last;
}
