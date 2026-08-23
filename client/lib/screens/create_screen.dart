import 'package:flutter/material.dart';

import 'package:client/api/create.dart';
import 'package:client/api/image.dart';
import 'package:client/api/update.dart';
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
  late TextEditingController _controller;
  late TextEditingController _filtered;
  bool _textLoading = false;
  bool _imageLoading = false;
  String _id = '';

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController();
    _filtered = TextEditingController();
  }

  @override
  void dispose() {
    _controller.dispose();
    _filtered.dispose();
    super.dispose();
  }

  Future<void> _fetchImage(String id) async {
    setState(() => _imageLoading = true);
    try {
      final image = await apiImage(id);

      setState(() {
        widget.state.player.flavor.image = image;
        _id = id;
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
      if (widget.state.player.id != '') {
        filtered = await apiUpdate(widget.state.player.id, _controller.text);
      } else {
        filtered = await apiCreate(_controller.text);
      }
      if (filtered['flavor']['error'] == '') {
        final f = Flavor();

        f.raw = _controller.text;
        f.day = widget.state.day;

        f.name = filtered['flavor']['name'];
        f.race = filtered['flavor']['race'];
        f.job = filtered['flavor']['job'];
        f.filtered = filtered['flavor']['text'];

        setState(() {
          widget.state.player.flavor = f;
          _filtered.text = widget.state.player.flavor.formatText();
        });

        _fetchImage(filtered['id']);
      } else {
        setState(() {
          _filtered.text = 'エラー:\n${filtered['flavor']['error']}';
        });
      }
    } catch (e) {
      // TODO
    } finally {
      setState(() => _textLoading = false);
    }
  }

  void _commit() {
    setState(() {
      final player = widget.state.player;
      player.flavors.add(player.flavor.clone());

      if (player.isForeigner) {
        player.inhabit = Inhabit.inhabitant;
        player.settled = widget.state.day;
        player.id = _id;
        widget.state.inhabitants.add(player);
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
                        controller: _controller,
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
                          (_controller.text.isEmpty ||
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
                        controller: _filtered,
                        maxLines: null,
                        readOnly: true,
                      ),
                    ),

                    const SizedBox(height: 24),

                    if (_imageLoading)
                      const CircularProgressIndicator()
                    else if (widget.state.player.flavor.image != null)
                      Center(
                        child: ConstrainedBox(
                          constraints: const BoxConstraints(maxWidth: 300),
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(12),
                            child: Image.memory(
                              widget.state.player.flavor.image!,
                            ),
                          ),
                        ),
                      ),

                    const SizedBox(height: 48),

                    ElevatedButton(
                      onPressed: () {
                        _commit();
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
