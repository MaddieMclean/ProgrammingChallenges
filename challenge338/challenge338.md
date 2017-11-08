<div class="md"><h1>Description</h1>

<p>Our maze explorer has some wierd rules for finding the exit and we are going to tell him if it is possible with his rules to get out.</p>

<p>Our explorer has the following rules:</p>

<ul>
<li>I always walk 6 blocks straight on and then turn 180° and start walking 6 blocks again</li>
<li>If a wall is in my way I turn to the right, if that not possible I turn to the left and if that is not possible I turn back from where I came.</li>
</ul>

<h1>Formal Inputs &amp; Outputs</h1>

<h2>Input description</h2>

<p>A maze with our explorer and the exit to reach</p>

<p>Legend: </p>

<pre><code>&gt; : Explorer looking East
&lt; : Explorer looking West
^ : Explorer looking North
v : Explorer looking south
E : Exit
# : wall
  : Clear passage way (empty space)
</code></pre>

<h3>Maze 1</h3>

<pre><code>#######
#&gt;   E#
#######
</code></pre>

<h3>Maze 2</h3>

<pre><code>#####E#
#&lt;    #
#######
</code></pre>

<h3>Maze 3</h3>

<pre><code>##########
#&gt;      E#
##########
</code></pre>

<h3>Maze 4</h3>

<pre><code>#####E#
##### #
#&gt;    #
##### #
#######
</code></pre>

<h3>Maze 5</h3>

<pre><code>#####E#
##### #
##### #
##### #
##### #
#&gt;    #
##### #
#######
</code></pre>

<h3>Challenge Maze</h3>

<pre><code>#########
#E ######
##      #
##### # #
#&gt;    ###
##### ###
##### ###
##### ###
##### ###
##### ###
##### ###
######### 
</code></pre>

<h3>Challenge Maze 2</h3>

<pre><code>#########
#E ######
##      #
## ## # #
##### # #
#&gt;    ###
##### ###
##### ###
##### ###
##### ###
##### ###
##### ###
######### 
</code></pre>

<h2>Output description</h2>

<p>Whetter it is possible to exit the maze </p>

<h3>Maze 1</h3>

<pre><code>possible/true/yay
</code></pre>

<h3>Maze 2</h3>

<pre><code>possible/true/yay
</code></pre>

<h3>Maze 3</h3>

<pre><code>impossible/not true/darn it
</code></pre>

<h3>Maze 4</h3>

<pre><code>possible/true/yay
</code></pre>

<h3>Maze 5</h3>

<pre><code>impossible/not true/darn it
</code></pre>

<h1>Notes/Hints</h1>

<p>Making a turn does not count as a step</p>

<h1>Several bonuses</h1>

<h2>Bonus 1:</h2>

<p>Generate your own (possible) maze.</p>

<h2>Bonus 2:</h2>

<p>Animate it and make a nice gif out off it.</p>

<h2>Bonus 3:</h2>

<p>Be the little voice in the head:</p>

<p>Instead of turning each 6 steps, you should implement a way to not turn if that would means that you can make it to the exit.</p>

<h1>Finally</h1>

<p>Have a good challenge idea?</p>

<p>Consider submitting it to <a href="/r/dailyprogrammer_ideas">/r/dailyprogrammer_ideas</a></p>
</div>