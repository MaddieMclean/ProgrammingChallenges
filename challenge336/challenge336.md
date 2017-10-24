<div class="md"><h1>Description</h1>

<p>Imagine a given set of numbers wherein some are cannibals. We define a cannibal as a larger number can eat a smaller number and <strong>increase its value by 1</strong>. There are no restrictions on how many numbers any given number can consume.   A number which has been consumed is <strong>no longer available</strong>.  </p>

<p>Your task is to determine the number of numbers which can have a value equal to or greater than a specified value.  </p>

<h1>Input Description</h1>

<p>You'll be given two integers, <em>i</em> and <em>j</em>, on the first line. <em>i</em> indicates how many values you'll be given, and <em>j</em> indicates the number of queries.  </p>

<p>Example:  </p>

<pre><code> 7 2     
 21 9 5 8 10 1 3
 10 15   
</code></pre>

<p>Based on the above description, 7 is number of values that you will be given.  2 is the number of queries.  </p>

<p>That means -<br>
* Query 1 - How many numbers can have the value of at least 10<br>
* Query 2 - How many numbers can have the value of at least 15</p>

<h1>Output Description</h1>

<p>Your program should calculate and show the number of numbers which are equal to or greater than the desired number.  For the sample input given, this will be - </p>

<pre><code> 4 2  
</code></pre>

<h2>Explanation</h2>

<p>For Query 1 - </p>

<p>The number 9 can consume the numbers 5 to raise its value to 10 </p>

<p>The number 8 can consume the numbers 1 and 3 to raise its value to 10.  </p>

<p>So including 21 and 10, we can get <strong>four</strong> numbers which have a value of at least 10.     </p>

<p>For Query 2 -  </p>

<p>The number 10 can consume the numbers 9,8,5,3, and 1 to raise its value to 15.  </p>

<p>So including 21, we can get <strong>two</strong> numbers which have a value of at least 15.  </p>

<h1>Credit</h1>

<p>This challenge was suggested by user <a href="/u/Lemvig42">/u/Lemvig42</a>, many thanks! If you have a challenge idea, please share it in <a href="/r/dailyprogrammer_ideas">/r/dailyprogrammer_ideas</a> and there's a good chance we'll use it</p>
</div>