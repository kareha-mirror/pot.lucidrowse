enum Inhabit {
  foreigner,
  inhabitant,
  hermit,
  forgotten;

  bool get isForeigner => this == Inhabit.foreigner;
}

class Flavor {
  String raw = '';
  String filtered = '';

  String name = '';
  String race = '';
  String job = '';

  String imageUrl = '';

  int day = -1;
}

class PlayerAction {
  String raw = '';
  String filtered = '';

  int day = -1;

  bool committed = false;
}

class Player {
  Inhabit inhabit = Inhabit.foreigner;
  int settled = -1;

  List<Flavor> flavors = [];
  Flavor flavor = Flavor();

  List<PlayerAction> actions = [];
  PlayerAction action = PlayerAction();

  String summary = '';
}
