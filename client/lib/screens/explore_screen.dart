import 'package:flutter/material.dart';

import 'package:client/api/list_players.dart';
import 'package:client/api/override_player.dart';
import 'package:client/api/region_state.dart';
import 'package:client/models/region.dart';
import 'package:client/screens/read_screen.dart';
import 'package:client/state.dart';
import 'package:client/utils/image_url.dart';
import 'package:client/widgets/region_card.dart';
import 'package:client/widgets/translucent_panel.dart';

class ExploreScreen extends StatefulWidget {
  final AppState state;

  const ExploreScreen({super.key, required this.state});

  @override
  State<ExploreScreen> createState() => _ExploreScreenState();
}

class _ExploreScreenState extends State<ExploreScreen> {
  int regionIndex = -1;
  List<dynamic> _players = [];
  String? _playersErrorMessage;
  late TextEditingController _stateController;
  String? _stateErrorMessage;

  @override
  void initState() {
    super.initState();

    _stateController = TextEditingController();
  }

  @override
  void dispose() {
    _stateController.dispose();

    super.dispose();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    _sync();
  }

  Future<void> _sync() async {
    if (await widget.state.sync()) {
      setState(() {});
    }
  }

  Future<void> _loadState() async {
    try {
      setState(() {
        _stateController.text = '';
        _stateErrorMessage = null;
      });

      final region = regions[regionIndex];
      final result = await apiRegionState(region.code);

      if (!mounted) return;

      setState(() => _stateController.text = result['state'] ?? '');
    } catch (e) {
      setState(() => _stateErrorMessage = e.toString());
    }
  }

  Future<void> _loadPlayers() async {
    try {
      setState(() {
        _players = [];
        _playersErrorMessage = null;
      });

      final region = regions[regionIndex];
      final result = await apiListPlayers(region.code);

      if (!mounted) return;

      setState(() => _players = result['players'] ?? []);
    } catch (e) {
      setState(() => _playersErrorMessage = e.toString());
    }
  }

  List<Widget> playersList() {
    List<Widget> list = [];
    for (final player in _players) {
      list.add(const SizedBox(height: 24));

      if (player['image-id'] != null) {
        list.add(
          ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 300),
            child: ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: Image.network(imageUrl(player['image-id'])),
            ),
          ),
        );
      }

      list.add(const SizedBox(height: 8));

      list.add(
        Wrap(
          spacing: 8,
          runSpacing: 8,
          crossAxisAlignment: WrapCrossAlignment.center,
          children: [
            TranslucentPanel(
              child: Text(
                '${player['name']} / ${player['race']} / ${player['job']}',
              ),
            ),

            diaryButton(context, widget.state, player['player-id']),

            if (widget.state.flavor == null && player['free'])
              TranslucentPanel(
                child: OutlinedButton(
                  onPressed: () async {
                    await apiOverridePlayer(player['player-id']);

                    widget.state.clear();
                    await _sync();

                    if (!mounted) return;

                    Navigator.pushNamed(context, '/home');
                  },
                  child: Text("入り込む"),
                ),
              ),
          ],
        ),
      );

      list.add(const SizedBox(height: 8));

      list.add(TranslucentPanel(child: Text(player['description'])));
    }
    return list;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          if (regionIndex == -1)
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: Image(
                image: AssetImage(
                  widget.state.committed
                      ? 'assets/images/night.webp'
                      : 'assets/images/home.webp',
                ),
                fit: BoxFit.cover,
              ),
            )
          else
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: Image(
                image: AssetImage(regions[regionIndex].imagePath),
                fit: BoxFit.cover,
              ),
            ),

          if (regionIndex == -1)
            SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
              child: Center(
                child: Column(
                  children: [
                    const SizedBox(height: 48),

                    TranslucentPanel(child: const Text('夢の世界の地域たち')),

                    const SizedBox(height: 12),

                    for (int index = 0; index < regions.length; index++)
                      Padding(
                        padding: EdgeInsets.symmetric(
                          horizontal: MediaQuery.sizeOf(context).width < 600
                              ? 6
                              : 48,
                          vertical: MediaQuery.sizeOf(context).width < 600
                              ? 3
                              : 24,
                        ),
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 600),
                          child: RegionCard(
                            region: regions[index],
                            onTap: () {
                              setState(() => regionIndex = index);
                              _loadState();
                              _loadPlayers();
                            },
                          ),
                        ),
                      ),

                    ElevatedButton(
                      onPressed: () => Navigator.pushNamed(context, '/home'),
                      child: const Text('ふたを閉じる'),
                    ),

                    const SizedBox(height: 96),
                  ],
                ),
              ),
            )
          else
            SingleChildScrollView(
              padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
              child: Center(
                child: Column(
                  children: [
                    const SizedBox(height: 48),

                    TranslucentPanel(child: const Text('夢の世界の住人たち')),

                    const SizedBox(height: 24),

                    if (_playersErrorMessage != null)
                      const SizedBox(height: 12),
                    if (_playersErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _playersErrorMessage!,
                          style: TextStyle(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    if (_players.isEmpty)
                      TranslucentPanel(child: const Text('まだ誰も住んでない。'))
                    else
                      ...playersList(),

                    const SizedBox(height: 48),

                    TranslucentPanel(child: const Text('地域の様子')),

                    const SizedBox(height: 12),

                    TranslucentPanel(
                      child: TextField(
                        controller: _stateController,
                        maxLines: null,
                        readOnly: true,
                      ),
                    ),

                    if (_stateErrorMessage != null) const SizedBox(height: 12),
                    if (_stateErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _stateErrorMessage!,
                          style: TextStyle(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    const SizedBox(height: 96),

                    ElevatedButton(
                      onPressed: () => setState(() => regionIndex = -1),
                      child: const Text('中を見直す'),
                    ),

                    const SizedBox(height: 96),
                  ],
                ),
              ),
            ),
        ],
      ),
    );
  }
}
