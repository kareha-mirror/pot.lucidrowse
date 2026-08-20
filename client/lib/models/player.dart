enum Inhabit {
  foreigner,
  inhabitant,
  hermit,
  forgotten;

  bool get isForeigner => this == Inhabit.foreigner;
}

class Flavor {
  String raw;
  String filtered;

  String name;
  String race;
  String job;

  String imageUrl;

  int time;

  Flavor.blank()
    : raw = '',
      filtered = '',
      name = '',
      race = '',
      job = '',
      imageUrl = '',
      time = 0;
}

class Action {
  String raw;
  String filtered;

  int time;

  Action.blank() : raw = '', filtered = '', time = 0;
}

class Player {
  Inhabit inhabit;

  List<Flavor> flavors;
  Flavor flavor;

  List<Action> actions;
  Action action;

  String summary;

  Player.unnamed()
    : inhabit = Inhabit.foreigner,
      flavors = [],
      flavor = Flavor.blank(),
      actions = [],
      action = Action.blank(),
      summary = '';
}
