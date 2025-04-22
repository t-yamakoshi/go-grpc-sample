module.exports = {
    disableEmoji: false,
    format: '{type}: {subject}',
    list: [
      'fix',
      'hotfix',
      'feat',
      'update',
      'refactor',
      'clean',
      'disable',
      'revert',
      'upgrade',
      'docs',
      'style',
      'build',
      'perf',
      'a11y',
      'test',
      'release',
      'breaking',
    ],
    maxMessageLength: 122,
    questions: ['type', 'subject', 'body', 'issues'],
    scopes: [],
    types: {
      fix: {
        description: 'バグ修正',
        emoji: '🐛',
        value: 'fix',
      },
      hotfix: {
        description: 'クリティカルなバグ修正',
        emoji: '🚑',
        value: 'hotfix',
      },
      feat: {
        description: '新機能の追加',
        emoji: '✨',
        value: 'feat',
      },
      update: {
        description: 'バグ修正ではない機能追加 (仕様変更も含む)',
        emoji: '🎉',
        value: 'update',
      },
      refactor: {
        description: '機能への変更・追加が伴わない変更 (リファクタリング)',
        emoji: '🔨',
        value: 'refactor',
      },
      clean: {
        description: '不要ファイルの削除、整理',
        emoji: '🔥',
        value: 'clean',
      },
      disable: {
        description: '機能の無効化 (コメントアウト等)',
        emoji: '🚫',
        value: 'disable',
      },
      revert: {
        description: '変更内容の取り消し',
        emoji: '⏪',
        value: 'revert',
      },
      upgrade: {
        description: 'ライブラリのアップデート',
        emoji: '🆙',
        value: 'upgrade',
      },
      docs: {
        description: 'ドキュメントの追加・修正',
        emoji: '📚',
        value: 'docs',
      },
      style: {
        description: 'コードのスタイル変更',
        emoji: '💄',
        value: 'style',
      },
      build: {
        description:
          'ビルドプロセスの変更、またはドキュメント生成など補助ツールに関連する変更',
        emoji: '🤖',
        value: 'build',
      },
      perf: {
        description: 'パフォーマンス改善',
        emoji: '⚡️',
        value: 'perf',
      },
      a11y: {
        description: 'アクセシビリティ改善',
        emoji: '👌',
        value: 'a11y',
      },
      test: {
        description: 'テストの追加・修正',
        emoji: '🚨',
        value: 'test',
      },
      release: {
        description: 'リリース作業に関する変更',
        emoji: '🚀',
        value: 'release',
      },
      breaking: {
        description: '以前のバージョンと互換性がない大規模な変更',
        emoji: '💥',
        value: 'BREAKING CHANGE',
      },
    },
    messages: {
      type: '変更の種類を選択してください:',
      subject: '変更内容を簡潔に記載してください:\n',
      body: '変更内容の詳細を記載してください (必須ではありません):\n',
      issues:
        '関連する Issue または Pull Request 番号 (e.g. #123) を記載してください (必須ではありません):\n',
    },
  };
  