/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('tableOfContents', [{
    'id': 'mechanics',
    'title': 'Passeando',
    'description': '<p>Bem-vindo ao passeio pela <a href="http://golang.org">linguagem de programação Go</a>. O passeio cobre as características mais importantes da linguagem, principalmente:</p>',
    'lessons': ['welcome']
}, {
    'id': 'basics',
    'title': 'Básico',
    'description': '<p>O ponto inicial é aprender todo básico da linguagem.</p><p>Declarando variáveis, chamando funções, e todas as coisas que você precisa saber para avançar para próxima lição.</p>',
    'lessons': ['basics', 'flowcontrol', 'moretypes']
}, {
    'id': 'methods',
    'title': 'Métodos e interfaces',
    'description': '<p>Aprenda como definir métodos nos tipos, como declarar interfaces e como juntar todas as coisas.</p>',
    'lessons': ['methods']
}, {
    'id': 'concurrency',
    'title': 'Concorrência',
    'description': '<p>Go oferece recursos de concorrência como parte do núcleo da linguagem.</p><p>Este módulo será sobre goroutines e canais, e como eles são usados para implementar diferentes padrões de concorrência.</p>',
    'lessons': ['concurrency']
}]).

// translation
value('translation', {
    'off': 'desligada',
    'on': 'ligada',
    'syntax': 'Mostrar Sintaxe',
    'lineno': 'Número nas linhas',
    'reset': 'Reset Slide',
    'format': 'Formate Código-Fonte',
    'kill': 'Mate o Programa',
    'run': 'Execute',
    'compile': 'Compile e Execute',
    'more': 'Opções',
    'toc': 'Conteúdo',
    'prev': 'Anterior',
    'next': 'Próximo',
    'waiting': 'Aguardando pelo servidor remoto...',
    'errcomm': 'Erro ao comunicar com servidor remoto.',
}).

// Config for codemirror plugin
value('ui.config', {
    codemirror: {
        mode: 'text/x-go',
        matchBrackets: true,
        lineNumbers: true,
        autofocus: true,
        indentWithTabs: true,
        indentUnit: 4,
        tabSize: 4,
        lineWrapping: true,
        extraKeys: {
            'Shift-Enter': function() {
                $('#run').click();
            },
            'Ctrl-Enter': function() {
                $('#format').click();
            },
            'PageDown': function() {
                return false;
            },
            'PageUp': function() {
                return false;
            },
        },
        // TODO: is there a better way to do this?
        // AngularJS values can't depend on factories.
        onChange: function() {
            if (window.codeChanged !== null) window.codeChanged();
        }
    }
});
