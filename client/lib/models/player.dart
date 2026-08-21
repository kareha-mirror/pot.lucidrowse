enum Inhabit { foreigner, inhabitant, hermit, forgotten }

class Flavor {
  String raw = '';
  String filtered = '';

  String name = '';
  String race = '';
  String job = '';

  String imageUrl = '';

  int day = 0;

  bool get hasFiltered => filtered != '';

  Flavor clone() {
    final flavor = Flavor();
    flavor.raw = raw;
    flavor.filtered = filtered;
    flavor.name = name;
    flavor.race = race;
    flavor.job = job;
    flavor.imageUrl = imageUrl;
    flavor.day = day;
    return flavor;
  }
}

class PlayerAction {
  String raw = '';
  String filtered = '';

  int day = 0;

  bool get hasFiltered => filtered != '';
}

class Player {
  Inhabit inhabit = Inhabit.foreigner;
  int settled = 0;

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
