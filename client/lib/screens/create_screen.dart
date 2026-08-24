import 'package:flutter/material.dart';

import 'package:client/api/commit_flavor.dart';
import 'package:client/api/image_flavor.dart';
import 'package:client/api/new_flavor.dart';
import 'package:client/api/new_player.dart';
import 'package:client/api/update_flavor.dart';
import 'package:client/constants.dart';
import 'package:client/models/player.dart';
import 'package:client/state/app_state.dart';
import 'package:client/widgets/translucent_panel.dart';

class CreateScreen extends StatefulWidget {
  final AppState state;

  const CreateScreen({super.key, required this.state});

  @override
  State<CreateScreen> createState() => _CreateScreenState();
}

class _CreateScreenState extends State<CreateScreen> {
  late TextEditingController _inputController;
  late TextEditingController _outputController;
  bool _textLoading = false;
  bool _imageLoading = false;

  @override
  void initState() {
    super.initState();
    _inputController = TextEditingController();
    _outputController = TextEditingController();
    if (widget.state.player.flavor.hasFiltered) {
      _outputController.text = widget.state.player.flavor.formatText();
    }
  }

  @override
  void dispose() {
    _inputController.dispose();
    _outputController.dispose();
    super.dispose();
  }

  Future<void> _fetchImage(String id) async {
    setState(() => _imageLoading = true);
    try {
      final image = await apiImageFlavor(id);

      setState(() {
        widget.state.player.flavor.image = image['image-id'];
      });
    } catch (e) {
      // TODO
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _fetchFlavor() async {
    setState(() => _textLoading = true);
    try {
      final Map<String, dynamic> filtered;
      if (widget.state.player.flavor.hasFiltered) {
        filtered = await apiUpdateFlavor(
          widget.state.player.id,
          _inputController.text,
        );
      } else {
        if (widget.state.player.id == '') {
          final created = await apiNewPlayer();
          widget.state.player.id = created['player-id'];
        }

        filtered = await apiNewFlavor(
          widget.state.player.id,
          _inputController.text,
        );
      }
      if (filtered['error'] == '') {
        final f = Flavor();

        f.raw = _inputController.text;
        f.day = widget.state.day;

        f.name = filtered['name'];
        f.race = filtered['race'];
        f.job = filtered['job'];
        f.filtered = filtered['description'];

        setState(() {
          widget.state.player.flavor = f;
          _outputController.text = widget.state.player.flavor.formatText();
        });

        _fetchImage(widget.state.player.id);
      } else {
        setState(() {
          _outputController.text = 'エラー:\n${filtered['error']}';
        });
      }
    } catch (e) {
      // TODO
    } finally {
      setState(() => _textLoading = false);
    }
  }

  Future<void> _commit() async {
    await apiCommitFlavor(widget.state.player.id);

    setState(() {
      final player = widget.state.player;
      player.flavors.add(player.flavor.clone());

      if (player.isForeigner) {
        player.inhabit = Inhabit.inhabitant;
        player.settled = widget.state.day;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          SizedBox(
            width: double.infinity,
            height: double.infinity,
            child: Image(
              image: AssetImage(
                widget.state.player.committed
                    ? 'assets/images/night.webp'
                    : 'assets/images/home.webp',
              ),
              fit: BoxFit.cover,
            ),
          ),

          SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Center(
              child: Column(
                children: [
                  if (widget.state.player.isForeigner) ...[
                    TranslucentPanel(
                      child: Column(
                        children: [
                          const Text('水晶球が心に呼びかける。'),
                          const Text('夢の世界へいざないましょう。'),
                        ],
                      ),
                    ),

                    const SizedBox(height: 24),

                    TranslucentPanel(
                      child: Column(
                        children: [
                          const Text('あなたはこの世界に自身の分身を作り出し、夢と関わっていくことになります。'),
                          const Text('この世界で夢見るあなたは、どんな人ですか、どんな生きものですか。'),
                        ],
                      ),
                    ),
                  ] else ...[
                    TranslucentPanel(
                      child: Column(
                        children: [
                          const Text('水晶球が心に呼びかける。'),
                          const Text('何か変化はありましたか。'),
                        ],
                      ),
                    ),

                    const SizedBox(height: 24),

                    TranslucentPanel(
                      child: const Text('あなたの分身に何か変化があるのなら、ここで新たに念じてください。'),
                    ),
                  ],

                  const SizedBox(height: 24),

                  Padding(
                    padding: EdgeInsets.all(
                      MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
                    ),
                    child: TranslucentPanel(
                      child: TextField(
                        controller: _inputController,
                        decoration: InputDecoration(
                          hintText: widget.state.player.isForeigner
                              ? '(どんな存在になりたいか、ここに念じよう。)'
                              : '(どんな変化があったか、ここに念じよう。)',
                        ),
                        onChanged: (String value) => setState(() {}),
                        maxLines: null,
                        maxLength: 140,
                        readOnly: _textLoading || _imageLoading,
                      ),
                    ),
                  ),

                  const SizedBox(height: 24),

                  TranslucentPanel(
                    child: ElevatedButton(
                      onPressed:
                          (_inputController.text.isEmpty ||
                              _textLoading ||
                              _imageLoading)
                          ? null
                          : _fetchFlavor,
                      child: const Text('念じる'),
                    ),
                  ),

                  if (_textLoading)
                    const CircularProgressIndicator()
                  else if (widget.state.player.flavor.hasFiltered) ...[
                    const SizedBox(height: 48),

                    TranslucentPanel(
                      child: const Text('あなたの言葉は夢に映され、こうなりました。'),
                    ),

                    const SizedBox(height: 24),

                    TranslucentPanel(
                      child: TextField(
                        controller: _outputController,
                        maxLines: null,
                        readOnly: true,
                      ),
                    ),

                    const SizedBox(height: 24),

                    if (_imageLoading)
                      const Column(
                        children: [
                          TranslucentPanel(child: const Text('姿を映し出しています…')),
                          const CircularProgressIndicator(),
                        ],
                      )
                    else if (widget.state.player.flavor.image != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.network(
                              '$apiBaseUrl/image/${widget.state.player.flavor.image!}',
                            ),
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed: () async {
                        await _commit();
                        Navigator.pop(context);
                      },
                      child: const Text('これでよし'),
                    ),
                  ],

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: () => Navigator.pop(context),
                    child: Text('目をそらす'),
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
