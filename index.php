<!DOCTYPE html>
<html lang="en">
    <head>
        <meta content="text/html" charset="utf-8" />
        <title>Sandbox (PHP)</title>
    </head>
    <body>
        <form method="post" action="https://webdevbasics.net/scripts/demo.php" id="opportunities">
        <?php
            function console_log($output, $with_script_tags = true) {
                # https://stackify.com/how-to-log-to-console-in-php
                $js_code = 'console.log(' . json_encode($output, JSON_HEX_TAG) . ');';
                
                if ($with_script_tags) {
                    $js_code = '<script>' . $js_code . '</script>';
                }
        
                echo $js_code;
            }

            for($i=0;$i<=10;$i++) {
                echo $i;
            }

            phpinfo(INFO_GENERAL);

            xdebug_info();
        ?>
    </body>
</html>