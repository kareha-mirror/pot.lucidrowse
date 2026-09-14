import 'package:flutter/material.dart';

import 'package:client/api/api.dart';
import 'package:client/screens/help_screen.dart';
import 'package:client/state/app_state.dart';
import 'package:client/widgets/password_field.dart';
import 'package:client/widgets/translucent_panel.dart';

class KeyScreen extends StatefulWidget {
  final AppState state;

  const KeyScreen({super.key, required this.state});

  @override
  State<KeyScreen> createState() => _KeyScreenState();
}

enum Menu { top, login, createKey, changePassword }

class _KeyScreenState extends State<KeyScreen> {
  Menu _menu = .top;

  late TextEditingController _usernameController;
  late TextEditingController _passwordController;
  late TextEditingController _newPasswordController;
  late TextEditingController _confirmController;

  @override
  void initState() {
    super.initState();

    _usernameController = .new();
    _passwordController = .new();
    _newPasswordController = .new();
    _confirmController = .new();
  }

  @override
  void dispose() {
    _usernameController.dispose();
    _passwordController.dispose();
    _newPasswordController.dispose();
    _confirmController.dispose();

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

  Future<void> _doLogin() async {
    try {
      await api.login(_usernameController.text, _passwordController.text);
      widget.state.clear();
      await _sync();

      if (!mounted) return;

      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const .new(content: Text('元の夢に帰って来ました。')));

      Navigator.pushNamed(context, '/');
    } catch (e) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(.new(content: Text(e.toString())));
    }
  }

  Future<void> _doCreateKey() async {
    if (_newPasswordController.text != _confirmController.text) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const .new(content: Text('確認の合い言葉が一致しません。')));
      return;
    }
    try {
      await api.ensureSession();
      await api.createKey(
        _usernameController.text,
        _newPasswordController.text,
      );
      widget.state.clear();
      await _sync();

      if (!mounted) return;

      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const .new(content: Text('秘密の鍵を作りました。')));

      Navigator.pushNamed(context, '/');
    } catch (e) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(.new(content: Text(e.toString())));
    }
  }

  Future<void> _doChangePassword() async {
    if (_newPasswordController.text != _confirmController.text) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const .new(content: Text('確認の合い言葉が一致しません。')));
      return;
    }
    try {
      await api.changePassword(
        _passwordController.text,
        _newPasswordController.text,
      );

      if (!mounted) return;

      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const .new(content: Text('合い言葉を変えました。')));

      Navigator.pushNamed(context, '/');
    } catch (e) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(.new(content: Text(e.toString())));
    }
  }

  Future<void> _doLogout() async {
    try {
      await api.logout();
      widget.state.clear();
      await _sync();

      if (!mounted) return;

      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const .new(content: Text('夢から覚めました。')));

      Navigator.pushNamed(context, '/');
    } catch (e) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(.new(content: Text(e.toString())));
    }
  }

  Widget _buildMenu() {
    return switch (_menu) {
      .top => _buildTop(),
      .login => _buildLogin(),
      .createKey => _buildCreateKey(),
      .changePassword => _buildChangePassword(),
    };
  }

  Widget _buildTop() {
    if (widget.state.user.name == null) {
      return Center(
        child: Column(
          children: [
            const SizedBox(height: 48),

            ElevatedButton(
              onPressed: () => setState(() => _menu = .login),
              child: const Text('持っている秘密の鍵を使う'),
            ),

            const SizedBox(height: 48),

            ElevatedButton(
              onPressed: () => setState(() => _menu = .createKey),
              child: const Text('新しい秘密の鍵を作る'),
            ),

            const SizedBox(height: 48),

            helpButton(context, .key),

            const SizedBox(height: 96),

            ElevatedButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('鍵を置く'),
            ),

            const SizedBox(height: 96),
          ],
        ),
      );
    } else {
      return Center(
        child: Column(
          children: [
            const SizedBox(height: 48),

            TranslucentPanel(
              child: Text('鍵の名前: ${widget.state.user.name ?? ''}'),
            ),

            const SizedBox(height: 48),

            ElevatedButton(
              onPressed: () => setState(() => _menu = .changePassword),
              child: const Text('合い言葉を変える'),
            ),

            const SizedBox(height: 48),

            ElevatedButton(onPressed: _doLogout, child: const Text('夢から覚める')),

            const SizedBox(height: 48),

            helpButton(context, .key),

            const SizedBox(height: 96),

            ElevatedButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('鍵を置く'),
            ),

            const SizedBox(height: 96),
          ],
        ),
      );
    }
  }

  Widget _buildLogin() {
    return Center(
      child: Column(
        children: [
          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('鍵の名前')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: TextField(
              controller: _usernameController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('合い言葉')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: PasswordField(
              controller: _passwordController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(
            child: ElevatedButton(
              onPressed:
                  (_usernameController.text == '' ||
                      _passwordController.text == '')
                  ? null
                  : _doLogin,
              child: const Text('鍵を開ける'),
            ),
          ),

          const SizedBox(height: 48),

          helpButton(context, .key),

          const SizedBox(height: 96),

          ElevatedButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('鍵を置く'),
          ),

          const SizedBox(height: 96),
        ],
      ),
    );
  }

  Widget _buildCreateKey() {
    return Center(
      child: Column(
        children: [
          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('新しい鍵の名前')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: TextField(
              controller: _usernameController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('新しい合い言葉')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: PasswordField(
              controller: _newPasswordController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('合い言葉をもう一度')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: PasswordField(
              controller: _confirmController,
              onChanged: (String value) => setState(() {}),
            ),
          ),
          const SizedBox(height: 48),

          TranslucentPanel(
            child: ElevatedButton(
              onPressed:
                  (_usernameController.text == '' ||
                      _newPasswordController.text == '' ||
                      _confirmController.text == '')
                  ? null
                  : _doCreateKey,
              child: const Text('新しい鍵を作る'),
            ),
          ),

          const SizedBox(height: 48),

          helpButton(context, .key),

          const SizedBox(height: 96),

          ElevatedButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('鍵を置く'),
          ),

          const SizedBox(height: 96),
        ],
      ),
    );
  }

  Widget _buildChangePassword() {
    return Center(
      child: Column(
        children: [
          const SizedBox(height: 48),

          TranslucentPanel(child: Text('鍵の名前: ${widget.state.user.name}')),

          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('現在の合い言葉')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: PasswordField(
              controller: _passwordController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('新しい合い言葉')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: PasswordField(
              controller: _newPasswordController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(child: const Text('合い言葉をもう一度')),

          const SizedBox(height: 12),

          TranslucentPanel(
            child: PasswordField(
              controller: _confirmController,
              onChanged: (String value) => setState(() {}),
            ),
          ),

          const SizedBox(height: 48),

          TranslucentPanel(
            child: ElevatedButton(
              onPressed:
                  (_passwordController.text == '' ||
                      _newPasswordController.text == '' ||
                      _confirmController.text == '')
                  ? null
                  : _doChangePassword,
              child: const Text('合い言葉を変える'),
            ),
          ),

          const SizedBox(height: 48),

          helpButton(context, .key),

          const SizedBox(height: 96),

          ElevatedButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('鍵を置く'),
          ),

          const SizedBox(height: 96),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return PopScope(
      canPop: false,
      onPopInvokedWithResult: (didPop, result) {
        if (!didPop) {
          Navigator.pop(context);
        }
      },
      child: Scaffold(
        body: Stack(
          children: [
            SizedBox(
              width: double.infinity,
              height: double.infinity,
              child: const Image(
                image: AssetImage('assets/images/key.webp'),
                fit: .cover,
              ),
            ),

            SingleChildScrollView(
              padding: const EdgeInsets.all(24),
              child: _buildMenu(),
            ),
          ],
        ),
      ),
    );
  }
}
