<?php

require_once('secret.php');

function encrypt($data) {
  $iv = openssl_random_pseudo_bytes(16);
  $enc = openssl_encrypt($data, 'AES-128-CBC', ENCRYPT_KEY, OPENSSL_RAW_DATA, $iv);
  return $iv.$enc;
}

function decrypt($data) {

  $iv = substr($data, 0, 16);
  $encrypted = substr($data, $ivlength);
  $dec = openssl_decrypt($encrypted, 'AES-128-CBC', ENCRYPT_KEY, OPENSSL_RAW_DATA, $iv);

  return $dec;
}

header('Content-Type: application/json');

$data = json_decode(file_get_contents('php://input'), true);
$mode = $data['mode'];

if ($mode === 'encrypt') {
  $content = base64_decode($data['content'], true);
  if ($content === false) {
    echo json_encode(['error' => 'invalid content']);
    die();
  }

  $result = encrypt($content);
  if ($result === false) {
    echo json_encode(['error' => openssl_error_string()]);
    die();
  }

  echo json_encode(['result' => base64_encode($result)]);
} else if ($mode === 'decrypt') {
  $content = base64_decode($data['content']);
  if ($content === false) {
    echo json_encode(['error' => 'invalid content']);
    die();
  } else if (strlen($content) < 16) {
    echo json_encode(['error' => 'invalid content']);
    die();
  }

  $result = decrypt($content);
  if ($result === false) {
    echo json_encode(['error' => openssl_error_string()]);
    die();
  }

  echo json_encode(['result' => 'ok. TODO: return the result without the flag']);

} else if ($mode === 'getflag') {
  $result = encrypt(FLAG);
  echo json_encode(['result' => base64_encode($result)]);
}
