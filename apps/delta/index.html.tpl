{{define "head"}}
<link rel="stylesheet" href="ui.css">
<link rel="stylesheet" href="delta.css">
{{end}}

{{define "body"}}
<div id="startup-screen">
    <h1>Delta</h1>
    <button id="host-game">Host Game</button>
    <button id="connect-to-game">Connect to Game</button>
    <button id="accept-invitation">Accept Invitation</button>
</div>
<div id="game-screen" style="display: none;">
    <div id="game-board">
    </div>
</div>
{{end}}