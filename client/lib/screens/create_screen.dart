import 'package:flutter/material.dart';

import 'package:client/api/commit_flavor.dart';
import 'package:client/api/image_flavor.dart';
import 'package:client/api/new_flavor.dart';
import 'package:client/api/new_player.dart';
import 'package:client/api/update_flavor.dart';
import 'package:client/models/player.dart';
import 'package:client/state.dart';
import 'package:client/utils/image_url.dart';
import 'package:client/widgets/translucent_panel.dart';

class CreateScreen extends StatefulWidget {
  final AppState state;

  const CreateScreen({super.key, required this.state});

  @override
  State<CreateScreen> createState() => _CreateScreenState();
}

class _CreateScreenState extends State<CreateScreen> {
  late TextEditingController _inputController;
  String? _flavorErrorMessage;
  late TextEditingController _outputController;
  Flavor _flavor = Flavor();
  bool _flavorLoading = false;
  bool _imageLoading = false;
  String? _imageErrorMessage;

  @override
  void initState() {
    super.initState();

    _inputController = TextEditingController();
    _outputController = TextEditingController();

    _flavor = Flavor.copy(widget.state.player.flavor);
    if (_flavor.hasDescription) {
      _outputController.text = _flavor.formatText();
    }
  }

  @override
  void dispose() {
    _inputController.dispose();
    _outputController.dispose();

    super.dispose();
  }

  bool _disabled() {
    final player = widget.state.player;
    if (player.committed) {
      return true;
    }
    return false;
  }

  Future<void> _loadImage() async {
    setState(() {
      _imageLoading = true;
      _imageErrorMessage = null;
    });
    try {
      final result = await apiImageFlavor(widget.state.player.id);

      setState(() {
        _flavor.imageId = result['image-id'];
      });
    } catch (e) {
      setState(() => _imageErrorMessage = e.toString());
    } finally {
      setState(() {
        _imageLoading = false;
      });
    }
  }

  Future<void> _loadFlavor() async {
    setState(() {
      _flavorLoading = true;
      _flavorErrorMessage = null;
    });
    try {
      final Map<String, dynamic> result;
      if (widget.state.player.flavor.hasDescription) {
        result = await apiUpdateFlavor(
          widget.state.player.id,
          _inputController.text,
        );
      } else {
        if (widget.state.player.id == '') {
          final created = await apiNewPlayer();
          widget.state.player.id = created['player-id'];
        }

        result = await apiNewFlavor(
          widget.state.player.id,
          _inputController.text,
        );
      }
      if (result['error'] == '') {
        final flavor = Flavor.fromJson(result);

        flavor.input = _inputController.text;

        setState(() {
          _outputController.text = flavor.formatText();
          _flavor = flavor;
        });

        _loadImage();
      } else {
        setState(() {
          _outputController.text = 'エラー:\n${result['error']}';
        });
      }
    } catch (e) {
      setState(() => _flavorErrorMessage = e.toString());
    } finally {
      setState(() => _flavorLoading = false);
    }
  }

  Future<void> _commit() async {
    await apiCommitFlavor(widget.state.player.id);

    setState(() {
      final player = widget.state.player;

      if (!player.inhabitant) {
        player.inhabitant = true;
      }

      player.flavor = _flavor;
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
                  if (!widget.state.player.inhabitant) ...[
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
                  ] else if (!_disabled()) ...[
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

                  if (!_disabled())
                    Padding(
                      padding: EdgeInsets.all(
                        MediaQuery.sizeOf(context).width < 600 ? 12 : 48,
                      ),
                      child: TranslucentPanel(
                        child: TextField(
                          controller: _inputController,
                          decoration: InputDecoration(
                            hintText: !widget.state.player.inhabitant
                                ? '(どんな存在になりたいか、ここに念じよう。)'
                                : '(どんな変化があったか、ここに念じよう。)',
                          ),
                          onChanged: (String value) => setState(() {}),
                          maxLines: null,
                          maxLength: 140,
                          readOnly: _flavorLoading || _imageLoading,
                        ),
                      ),
                    ),

                  if (!_disabled()) const SizedBox(height: 24),

                  if (!_disabled())
                    TranslucentPanel(
                      child: ElevatedButton(
                        onPressed:
                            (_inputController.text.isEmpty ||
                                _flavorLoading ||
                                _imageLoading)
                            ? null
                            : _loadFlavor,
                        child: const Text('念じる'),
                      ),
                    ),

                  if (_flavorErrorMessage != null) const SizedBox(height: 12),
                  if (_flavorErrorMessage != null)
                    TranslucentPanel(
                      child: Text(
                        _flavorErrorMessage!,
                        style: TextStyle(
                          color: Theme.of(context).colorScheme.error,
                        ),
                      ),
                    ),

                  if (_flavorLoading)
                    const CircularProgressIndicator()
                  else if (_flavor.hasDescription) ...[
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
                          TranslucentPanel(child: Text('姿を映し出しています…')),
                          CircularProgressIndicator(),
                        ],
                      )
                    else if (_flavor.imageId != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.network(imageUrl(_flavor.imageId!)),
                          ),
                        ),
                      ),

                    if (_imageErrorMessage != null) const SizedBox(height: 12),
                    if (_imageErrorMessage != null)
                      TranslucentPanel(
                        child: Text(
                          _imageErrorMessage!,
                          style: TextStyle(
                            color: Theme.of(context).colorScheme.error,
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed:
                          (_flavorLoading || _imageLoading || _disabled())
                          ? null
                          : () async {
                              await _commit();

                              if (!context.mounted) return;

                              Navigator.pop(context);
                            },
                      child: const Text('これでよし'),
                    ),
                  ],

                  const SizedBox(height: 96),

                  ElevatedButton(
                    onPressed: (_flavorLoading || _imageLoading)
                        ? null
                        : () => Navigator.pop(context),
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
